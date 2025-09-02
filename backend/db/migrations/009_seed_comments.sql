-- +goose Up
-- Insert comments on at least 7 posts
INSERT INTO comments (id, post_id, author_id, parent_id, body, created_at, updated_at) VALUES

-- Comments on "The Future of Web Development in 2024" by Sam
('c1936d7d-4f8a-7dd0-9f3b-101112131415', '01936d7d-4f8a-7dd0-9f3b-101112131415', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', NULL, 
'Great insights Sam! AI tools have definitely changed my development workflow. The productivity gains are remarkable.', 
NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),

('c1936d7d-4f8a-7dd0-9f3b-101112131416', '01936d7d-4f8a-7dd0-9f3b-101112131415', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', NULL, 
'WebAssembly is fascinating! I am curious about the learning curve for developers coming from traditional JS backgrounds.',
NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),

-- Reply to Emma's comment
('c1936d7d-4f8a-7dd0-9f3b-101112131417', '01936d7d-4f8a-7dd0-9f3b-101112131415', '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', 'c1936d7d-4f8a-7dd0-9f3b-101112131416',
'@emma The learning curve is moderate if you already know C/C++ or Rust. The tooling has improved significantly!',
NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),

-- Comments on "Understanding Large Language Models" by Maya
('c1936d7d-4f8a-7dd0-9f3b-202122232425', '01936d7d-4f8a-7dd0-9f3b-202122232425', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', NULL,
'Excellent breakdown of transformer architecture! This helped me understand why attention mechanisms are so powerful.',
NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),

('c1936d7d-4f8a-7dd0-9f3b-202122232426', '01936d7d-4f8a-7dd0-9f3b-202122232425', '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', NULL,
'The emergent abilities aspect is mind-blowing. It''s like the models discover new capabilities we never explicitly taught them.',
NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),

-- Comments on "From Idea to MVP" by Alex
('c1936d7d-4f8a-7dd0-9f3b-303132333435', '01936d7d-4f8a-7dd0-9f3b-303132333435', '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', NULL,
'The validation point hits home. I spent 6 months building something nobody wanted. Hard lesson learned!',
NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),

('c1936d7d-4f8a-7dd0-9f3b-303132333436', '01936d7d-4f8a-7dd0-9f3b-303132333435', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', NULL,
'Team chemistry is so underrated. Skills can be taught, but cultural fit is harder to fix.',
NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),

-- Comments on "Design Systems" by Emma
('c1936d7d-4f8a-7dd0-9f3b-404142434445', '01936d7d-4f8a-7dd0-9f3b-404142434445', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', NULL,
'We''re implementing a design system at our company. The governance part is the trickiest - how do you handle component updates?',
NOW() - INTERVAL '11 days', NOW() - INTERVAL '11 days'),

-- Reply to Maya's comment
('c1936d7d-4f8a-7dd0-9f3b-404142434446', '01936d7d-4f8a-7dd0-9f3b-404142434445', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', 'c1936d7d-4f8a-7dd0-9f3b-404142434445',
'@maya We use semantic versioning for components and maintain a changelog. Also regular office hours for questions!',
NOW() - INTERVAL '11 days', NOW() - INTERVAL '11 days'),

-- Comments on "Machine Learning in Production" by Jon
('c1936d7d-4f8a-7dd0-9f3b-505152535455', '01936d7d-4f8a-7dd0-9f3b-505152535455', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', NULL,
'Data drift is such a pain! Have you found any tools that work well for automated detection and alerting?',
NOW() - INTERVAL '12 days', NOW() - INTERVAL '12 days'),

-- Comments on "Mastering TypeScript" by Sam
('c1936d7d-4f8a-7dd0-9f3b-101112131418', '01936d7d-4f8a-7dd0-9f3b-101112131416', '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', NULL,
'Conditional types are game changers! They''ve helped me create much better APIs with compile-time safety.',
NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),

('c1936d7d-4f8a-7dd0-9f3b-101112131419', '01936d7d-4f8a-7dd0-9f3b-101112131416', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', NULL,
'Template literal types are so powerful for creating type-safe event systems. Great examples!',
NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),

-- Comments on "Building Remote Teams" by Alex
('c1936d7d-4f8a-7dd0-9f3b-303132333437', '01936d7d-4f8a-7dd0-9f3b-303132333437', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', NULL,
'The virtual coffee chat idea is brilliant! We''ve been struggling with team bonding remotely.',
NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),

('c1936d7d-4f8a-7dd0-9f3b-303132333438', '01936d7d-4f8a-7dd0-9f3b-303132333437', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', NULL,
'Results-oriented culture is key. Trust your team to deliver rather than micromanaging their hours.',
NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days');

-- +goose Down
DELETE FROM comments WHERE id IN (
  'c1936d7d-4f8a-7dd0-9f3b-101112131415',
  'c1936d7d-4f8a-7dd0-9f3b-101112131416',
  'c1936d7d-4f8a-7dd0-9f3b-101112131417',
  'c1936d7d-4f8a-7dd0-9f3b-202122232425',
  'c1936d7d-4f8a-7dd0-9f3b-202122232426',
  'c1936d7d-4f8a-7dd0-9f3b-303132333435',
  'c1936d7d-4f8a-7dd0-9f3b-303132333436',
  'c1936d7d-4f8a-7dd0-9f3b-404142434445',
  'c1936d7d-4f8a-7dd0-9f3b-404142434446',
  'c1936d7d-4f8a-7dd0-9f3b-505152535455',
  'c1936d7d-4f8a-7dd0-9f3b-101112131418',
  'c1936d7d-4f8a-7dd0-9f3b-101112131419',
  'c1936d7d-4f8a-7dd0-9f3b-303132333437',
  'c1936d7d-4f8a-7dd0-9f3b-303132333438'
);