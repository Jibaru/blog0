package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	ID          string          `sql:"id,primary"`
	AuthorID    string          `sql:"author_id"`
	Title       string          `sql:"title"`
	Slug        string          `sql:"slug"`
	RawMarkdown string          `sql:"raw_markdown"`
	Summary     string          `sql:"summary"`
	Tags        json.RawMessage `sql:"tags"`
	PublishedAt *time.Time      `sql:"published_at"`
	CreatedAt   time.Time       `sql:"created_at"`
	UpdatedAt   time.Time       `sql:"updated_at"`

	RawMarkdownAudioURL *string `sql:"raw_markdown_audio_url"`
	SummaryAudioURL     *string `sql:"summary_audio_url"`
}

func NewPost(id string, authorID string, title string, slug string, rawMarkdown string, summary string, tags []string) (*Post, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	if authorID == "" {
		return nil, fmt.Errorf("author ID cannot be empty")
	}

	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	if slug == "" {
		return nil, fmt.Errorf("slug cannot be empty")
	}

	if rawMarkdown == "" {
		return nil, fmt.Errorf("raw markdown cannot be empty")
	}

	if summary == "" {
		return nil, fmt.Errorf("summary cannot be empty")
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("tags cannot be empty")
	}

	rawTags, err := json.Marshal(tags)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tags: %w", err)
	}

	now := time.Now()
	return &Post{
		ID:          id,
		AuthorID:    authorID,
		Title:       title,
		Slug:        slug,
		RawMarkdown: rawMarkdown,
		Summary:     summary,
		Tags:        rawTags,
		PublishedAt: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func NewPublishedPost(id string, authorID string, title string, slug string, rawMarkdown string, summary string, tags []string, publishedAt time.Time) (*Post, error) {
	post, err := NewPost(id, authorID, title, slug, rawMarkdown, summary, tags)
	if err != nil {
		return nil, err
	}

	post.PublishedAt = &publishedAt
	return post, nil
}

func (p *Post) Publish(publishedAt time.Time) {
	p.PublishedAt = &publishedAt
	p.UpdatedAt = time.Now()
}

func (p *Post) Update(title string, slug string, rawMarkdown string, summary string, tags []string) error {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	if slug == "" {
		return fmt.Errorf("slug cannot be empty")
	}

	if rawMarkdown == "" {
		return fmt.Errorf("raw markdown cannot be empty")
	}

	if summary == "" {
		return fmt.Errorf("summary cannot be empty")
	}

	if len(tags) == 0 {
		return fmt.Errorf("tags cannot be empty")
	}

	rawTags, err := json.Marshal(tags)
	if err != nil {
		return fmt.Errorf("failed to marshal tags: %w", err)
	}

	p.Title = title
	p.Slug = slug
	p.RawMarkdown = rawMarkdown
	p.UpdatedAt = time.Now()
	p.Summary = summary
	p.Tags = rawTags
	return nil
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) ItsTags() []string {
	var tags []string
	if len(p.Tags) == 0 {
		return tags
	}
	_ = json.Unmarshal(p.Tags, &tags)
	return tags
}

type PostCreated struct {
	PostID string
}

type PostUpdated struct {
	PostID string
}
