import { logger, task } from "@trigger.dev/sdk/v3";
import { db } from "../db/index";
import { posts } from "../db/schema";
import { eq } from "drizzle-orm";
import { getElevenLabsClient } from "../lib/elevenlabs/client";
import { uploadStreamToUploadThing } from "../lib/uploadthing/utils";

export const generatePostAudioTask = task({
  id: "generate-post-audio",
  maxDuration: 5 * 60, // 5 min
  run: async (
    payload: {
      postId: string;
    },
    { ctx }
  ) => {
    logger.log("generate post audio started", { payload, ctx });

    const selectedPosts = await db
      .select()
      .from(posts)
      .where(eq(posts.id, payload.postId))
      .limit(1);

    if (!selectedPosts[0]) {
      throw new Error(`Post with id ${payload.postId} not found`);
    }

    if (!selectedPosts[0].raw_markdown) {
      throw new Error(`Post ${payload.postId} has no raw_markdown`);
    }

    const rawMarkdownAudioUrl = await generateAudioUrl(
      selectedPosts[0].raw_markdown,
      selectedPosts[0].id,
      `${selectedPosts[0].id}_raw_markdown}`
    );
    const summaryAudioUrl = await generateAudioUrl(
      selectedPosts[0].summary,
      selectedPosts[0].id,
      `${selectedPosts[0].id}_summary}`
    );

    await db
      .update(posts)
      .set({
        raw_markdown: rawMarkdownAudioUrl,
        summary_audio_url: summaryAudioUrl,
      })
      .where(eq(posts.id, payload.postId));

    logger.log("generate post audio completed", {
      postId: payload.postId,
      rawMarkdownAudioUrl,
      summaryAudioUrl,
    });

    return { rawMarkdownAudioUrl, summaryAudioUrl };
  },
});

async function generateAudioUrl(
  content: string,
  postId: string,
  name: string
): Promise<string> {
  const client = getElevenLabsClient();

  const audioStream = await client.textToSpeech.convert(
    "JBFqnCBsd6RMkjVDRZzb",
    {
      text: content,
      modelId: "eleven_multilingual_v2",
      outputFormat: "mp3_44100_128",
    }
  );

  return await uploadStreamToUploadThing(audioStream, `${name}.mp3`);
}
