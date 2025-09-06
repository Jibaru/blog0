import { UTApi } from "uploadthing/server";

const utapi = new UTApi();

export async function uploadStreamToUploadThing(
  stream: ReadableStream,
  filename: string
): Promise<string> {
  // Convert stream to buffer
  const response = new Response(stream);
  const buffer = await response.arrayBuffer();

  // Create a File object from the buffer
  const file = new File([buffer], filename, { type: "audio/mpeg" });

  // Upload using UTApi
  const uploadResult = await utapi.uploadFiles(file);

  if (uploadResult.error) {
    throw new Error(`UploadThing error: ${uploadResult.error.message}`);
  }

  return uploadResult.data.ufsUrl;
}
