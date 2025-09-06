import { logger, schedules } from "@trigger.dev/sdk/v3";
import Blog0ApiClient from "../lib/blog0/api-client";
import { generateObject } from "ai";
import z from "zod";
import { openai } from "@ai-sdk/openai";

export const generatePostTask = schedules.task({
  id: "generate-post",
  maxDuration: 5 * 60, // 5 min
  // cron every day at 00:00 UTC
  cron: "0 0 * * *",
  run: async (payload) => {
    logger.log("generate post started", { timestamp: payload.timestamp });

    try {
      const { title, slug, rawMarkdown } = await generateContent();

      const authorizationValue = process.env.PROCESSOR_PWD;

      const apiClient = new Blog0ApiClient({
        apiToken: authorizationValue,
      });

      const post = await apiClient.createPost({
        raw_markdown: rawMarkdown,
        slug: slug,
        title: title,
        publish: true,
      });

      logger.log("generate post completed", { post });
    } catch (error) {
      logger.error("generate post failed", { error });
    }
  },
});

async function generateContent(): Promise<{
  title: string;
  slug: string;
  rawMarkdown: string;
}> {
  const { object } = await generateObject({
    model: openai("gpt-4o"),
    schema: z.object({
      post: z.object({
        title: z.string(),
        markdownContent: z.string().describe("markdown content of the post"),
      }),
    }),
    prompt:
      "Generate a post title and content in markdown with original data about something of technology in the current year and month.",
  });

  return {
    title: object.post.title,
    rawMarkdown: object.post.markdownContent,
    slug: generateSlug(object.post.title),
  };
}

function generateSlug(title: string): string {
  // Normalize the title: lowercase, remove special chars, replace spaces with dashes
  const slugPart = title
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9\s-]/g, "") // remove non-alphanumeric chars (except spaces and dashes)
    .replace(/\s+/g, "-") // replace spaces with dashes
    .replace(/-+/g, "-"); // collapse multiple dashes

  // Get current Unix timestamp
  const timestamp = Math.floor(Date.now() / 1000);

  return `${slugPart}-${timestamp}`;
}
