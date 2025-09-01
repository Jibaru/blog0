-- +goose Up
-- FOLLOWS (follow other users)
CREATE TABLE follows (
  follower_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  followee_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL,
  PRIMARY KEY (follower_id, followee_id)
);

CREATE INDEX idx_follows_followee ON follows(followee_id);

-- +goose Down
DROP INDEX IF EXISTS idx_follows_followee;
DROP TABLE IF EXISTS follows;