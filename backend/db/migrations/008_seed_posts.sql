-- +goose Up
-- Insert 25 posts (5 per user)
INSERT INTO posts (id, author_id, title, slug, raw_markdown, published_at, created_at, updated_at) VALUES

-- Sam's posts (Tech Writer)  
('01936d7d-4f8a-7dd0-9f3b-101112131415', '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', 'The Future of Web Development in 2024', 'future-web-development-2024', 'The Future of Web Development in 2024

Web development continues to evolve rapidly. Key trends include AI-powered development tools, WebAssembly going mainstream, and serverless architecture evolution.', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),

('01936d7d-4f8a-7dd0-9f3b-101112131416', '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', 'Mastering TypeScript: Advanced Patterns', 'mastering-typescript-advanced-patterns', 'TypeScript has become the standard for large-scale JavaScript applications. Advanced patterns like conditional types and template literal types unlock powerful type-safe APIs.', NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),

('01936d7d-4f8a-7dd0-9f3b-101112131417', '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', 'Building Scalable APIs with Go and PostgreSQL', 'scalable-apis-go-postgresql', 'Go simplicity and performance make it excellent for robust APIs. Combined with PostgreSQL, you have a powerful stack for scalable applications.', NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),

('01936d7d-4f8a-7dd0-9f3b-101112131418', '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', 'Docker Best Practices for Development Teams', 'docker-best-practices-dev-teams', 'Docker has revolutionized application development and deployment. Essential practices include multi-stage builds, security considerations, and proper image management.', NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),

('01936d7d-4f8a-7dd0-9f3b-101112131419', '01936d7d-4f8a-7dd0-9f3b-1c2e3a4b5c6d', 'React Performance Optimization Techniques', 'react-performance-optimization', 'Performance is crucial for user experience. Proven techniques include memoization strategies, code splitting, and virtual scrolling for large lists.', NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),

-- Maya's posts (AI Researcher)
('01936d7d-4f8a-7dd0-9f3b-202122232425', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', 'Understanding Large Language Models: A Deep Dive', 'understanding-large-language-models', 'Large Language Models have transformed AI interaction. The transformer architecture and attention mechanism enable understanding context across long sequences.', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),

('01936d7d-4f8a-7dd0-9f3b-202122232426', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', 'The Ethics of AI: Navigating Bias and Fairness', 'ethics-ai-bias-fairness', 'As AI systems become prevalent, addressing ethical concerns becomes paramount. We must tackle bias types and implement mitigation strategies.', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),

('01936d7d-4f8a-7dd0-9f3b-202122232427', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', 'Prompt Engineering: The Art of Talking to AI', 'prompt-engineering-art-talking-ai', 'Effective prompt engineering is crucial for getting the best results from language models. Key principles include specificity, context, and iteration.', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),

('01936d7d-4f8a-7dd0-9f3b-202122232428', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', 'Computer Vision in 2024: Beyond Image Recognition', 'computer-vision-2024-beyond-recognition', 'Computer vision has evolved beyond image classification. Multimodal understanding and 3D scene understanding open unprecedented possibilities.', NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),

('01936d7d-4f8a-7dd0-9f3b-202122232429', '01936d7d-4f8a-7dd0-9f3b-2c3d4e5f6a7b', 'The Science of AI Safety: Alignment and Control', 'science-ai-safety-alignment-control', 'Ensuring AI safety and alignment with human values is critical as systems become more capable. Research focuses on constitutional AI and interpretability.', NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),

-- Alex's posts (Startup Founder)
('01936d7d-4f8a-7dd0-9f3b-303132333435', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', 'From Idea to MVP: A Founders Journey', 'idea-to-mvp-founders-journey', 'Building a startup is a rollercoaster. The biggest mistake is building without validation. Success comes from consistent execution, not brilliant ideas.', NOW() - INTERVAL '11 days', NOW() - INTERVAL '11 days', NOW() - INTERVAL '11 days'),

('01936d7d-4f8a-7dd0-9f3b-303132333436', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', 'Fundraising in a Down Market: Strategies That Work', 'fundraising-down-market-strategies', 'Raising capital in bear markets requires different strategies. Focus on fundamentals, build investor relations early, and consider alternative funding sources.', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),

('01936d7d-4f8a-7dd0-9f3b-303132333437', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', 'Building Remote Teams: Lessons from a Distributed Startup', 'building-remote-teams-distributed-startup', 'Being fully remote since day one taught us valuable lessons. Communication is everything, culture building requires intention, and results matter more than hours.', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),

('01936d7d-4f8a-7dd0-9f3b-303132333438', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', 'Product-Market Fit: How to Know When You Have Found It', 'product-market-fit-how-to-know', 'Product-market fit is the startup holy grail. Watch for quantitative signals like organic growth and qualitative signals like customer upset when service is down.', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),

('01936d7d-4f8a-7dd0-9f3b-303132333439', '01936d7d-4f8a-7dd0-9f3b-3d4e5f6a7b8c', 'Scaling Engineering Teams: When and How to Hire', 'scaling-engineering-teams-when-how-hire', 'Knowing when to scale engineering teams is crucial. Scale when backlog outpaces capacity and focus on cultural fit and learning ability over pure skills.', NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),

-- Emma's posts (Designer)
('01936d7d-4f8a-7dd0-9f3b-404142434445', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', 'Design Systems: Building Consistency at Scale', 'design-systems-consistency-at-scale', 'A well-crafted design system is the backbone of great product experiences. It includes foundation elements, component libraries, and governance processes.', NOW() - INTERVAL '12 days', NOW() - INTERVAL '12 days', NOW() - INTERVAL '12 days'),

('01936d7d-4f8a-7dd0-9f3b-404142434446', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', 'Accessibility in Design: Beyond Compliance', 'accessibility-design-beyond-compliance', 'Accessibility is about creating inclusive experiences for everyone. Universal design principles and practical implementation benefit all users, not just those with disabilities.', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),

('01936d7d-4f8a-7dd0-9f3b-404142434447', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', 'The Psychology of Color in Digital Interfaces', 'psychology-color-digital-interfaces', 'Color choices significantly impact user behavior and emotional responses. Consider cultural differences and accessibility when applying color psychology principles.', NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),

('01936d7d-4f8a-7dd0-9f3b-404142434448', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', 'User Research Methods: Choosing the Right Approach', 'user-research-methods-choosing-approach', 'Different research questions require different methods. Choose between exploratory research for insights and evaluative research for validation based on your goals.', NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),

('01936d7d-4f8a-7dd0-9f3b-404142434449', '01936d7d-4f8a-7dd0-9f3b-4e5f6a7b8c9d', 'Designing for Mobile-First: Beyond Responsive Design', 'designing-mobile-first-beyond-responsive', 'Mobile-first design is about understanding mobile user behavior, not just fitting content on smaller screens. Consider context, touch interactions, and performance.', NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),

-- Jon's posts (Data Scientist)
('01936d7d-4f8a-7dd0-9f3b-505152535455', '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', 'Machine Learning in Production: Lessons from the Trenches', 'machine-learning-production-lessons-trenches', 'Deploying ML models to production differs vastly from training in notebooks. Focus on model versioning, monitoring, and infrastructure considerations for scale.', NOW() - INTERVAL '13 days', NOW() - INTERVAL '13 days', NOW() - INTERVAL '13 days'),

('01936d7d-4f8a-7dd0-9f3b-505152535456', '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', 'Data Engineering for Machine Learning: Building Robust Pipelines', 'data-engineering-ml-robust-pipelines', 'Great ML models need great data pipelines. Build scalable infrastructure with event-driven processing, data quality checks, and proper monitoring.', NOW() - INTERVAL '11 days', NOW() - INTERVAL '11 days', NOW() - INTERVAL '11 days'),

('01936d7d-4f8a-7dd0-9f3b-505152535457', '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', 'A/B Testing for Data Scientists: Statistical Rigor Meets Practical Reality', 'ab-testing-data-scientists-statistical-rigor', 'A/B testing is more nuanced than it appears. Good experiments require careful planning, proper randomization, and statistical rigor to avoid common pitfalls.', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),

('01936d7d-4f8a-7dd0-9f3b-505152535458', '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', 'Feature Engineering: The Art of Creating Predictive Variables', 'feature-engineering-art-predictive-variables', 'Feature engineering can make or break ML models. Master techniques for numerical, categorical, and time series features while collaborating with domain experts.', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),

('01936d7d-4f8a-7dd0-9f3b-505152535459', '01936d7d-4f8a-7dd0-9f3b-5f6a7b8c9d0e', 'Interpreting Machine Learning Models: Beyond Black Boxes', 'interpreting-ml-models-beyond-black-boxes', 'Model interpretability builds trust and meets regulatory requirements. Use both global and local interpretability techniques to understand model decisions.', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days');

-- +goose Down
DELETE FROM posts WHERE id IN (
  '01936d7d-4f8a-7dd0-9f3b-101112131415',
  '01936d7d-4f8a-7dd0-9f3b-101112131416', 
  '01936d7d-4f8a-7dd0-9f3b-101112131417',
  '01936d7d-4f8a-7dd0-9f3b-101112131418',
  '01936d7d-4f8a-7dd0-9f3b-101112131419',
  '01936d7d-4f8a-7dd0-9f3b-202122232425',
  '01936d7d-4f8a-7dd0-9f3b-202122232426',
  '01936d7d-4f8a-7dd0-9f3b-202122232427',
  '01936d7d-4f8a-7dd0-9f3b-202122232428',
  '01936d7d-4f8a-7dd0-9f3b-202122232429',
  '01936d7d-4f8a-7dd0-9f3b-303132333435',
  '01936d7d-4f8a-7dd0-9f3b-303132333436',
  '01936d7d-4f8a-7dd0-9f3b-303132333437',
  '01936d7d-4f8a-7dd0-9f3b-303132333438',
  '01936d7d-4f8a-7dd0-9f3b-303132333439',
  '01936d7d-4f8a-7dd0-9f3b-404142434445',
  '01936d7d-4f8a-7dd0-9f3b-404142434446',
  '01936d7d-4f8a-7dd0-9f3b-404142434447',
  '01936d7d-4f8a-7dd0-9f3b-404142434448',
  '01936d7d-4f8a-7dd0-9f3b-404142434449',
  '01936d7d-4f8a-7dd0-9f3b-505152535455',
  '01936d7d-4f8a-7dd0-9f3b-505152535456',
  '01936d7d-4f8a-7dd0-9f3b-505152535457',
  '01936d7d-4f8a-7dd0-9f3b-505152535458',
  '01936d7d-4f8a-7dd0-9f3b-505152535459'
);