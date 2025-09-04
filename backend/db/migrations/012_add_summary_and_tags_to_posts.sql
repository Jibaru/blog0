-- +goose Up
ALTER TABLE posts
  ADD COLUMN summary TEXT,              -- short description or excerpt
  ADD COLUMN tags JSONB DEFAULT '[]';   -- array of strings (e.g. ["go", "db", "web"])

-- +goose Down
ALTER TABLE posts
  DROP COLUMN IF EXISTS tags,
  DROP COLUMN IF EXISTS summary;
