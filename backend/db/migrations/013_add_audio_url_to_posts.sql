-- +goose Up
ALTER TABLE posts
  ADD COLUMN raw_markdown_audio_url TEXT,
  ADD COLUMN summary_audio_url TEXT;

-- +goose Down
ALTER TABLE posts
  DROP COLUMN IF EXISTS raw_markdown_audio_url;
ALTER TABLE posts
  DROP COLUMN IF EXISTS summary_audio_url;
