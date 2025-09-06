import { pgTable, text, timestamp, uuid } from "drizzle-orm/pg-core";

export const posts = pgTable("posts", {
  id: uuid("id").primaryKey().notNull(),
  author_id: uuid("author_id").notNull(),
  title: text("title").notNull(),
  slug: text("slug"),
  raw_markdown: text("raw_markdown").notNull(),
  summary: text("summary").notNull(),
  raw_markdown_audio_url: text("raw_markdown_audio_url"),
  summary_audio_url: text("summary_audio_url"),
  published_at: timestamp("published_at", { withTimezone: true }),
  created_at: timestamp("created_at", { withTimezone: true }).notNull(),
  updated_at: timestamp("updated_at", { withTimezone: true }).notNull(),
});
