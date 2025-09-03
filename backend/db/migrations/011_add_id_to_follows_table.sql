-- +goose Up
-- Add ID column as primary key to follows table
ALTER TABLE follows DROP CONSTRAINT follows_pkey;
ALTER TABLE follows ADD COLUMN id UUID PRIMARY KEY DEFAULT gen_random_uuid();
CREATE UNIQUE INDEX idx_follows_follower_followee ON follows(follower_id, followee_id);

-- +goose Down
DROP INDEX IF EXISTS idx_follows_follower_followee;
ALTER TABLE follows DROP COLUMN id;
ALTER TABLE follows ADD PRIMARY KEY (follower_id, followee_id);