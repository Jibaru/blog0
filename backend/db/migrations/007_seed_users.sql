-- +goose Up
-- Insert 5 users
INSERT INTO users (id, username, email) VALUES
('01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', 'tech_writer_sam', 'sam@techblog.com'),
('01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', 'ai_researcher_maya', 'maya@airesearch.org'),
('01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', 'startup_founder_alex', 'alex@startuplife.io'),
('01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', 'designer_emma', 'emma@designstudio.co'),
('01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', 'data_scientist_jon', 'jon@datascience.net');

-- +goose Down
DELETE FROM users WHERE id IN (
  '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d',
  '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b',
  '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c',
  '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d',
  '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e'
);