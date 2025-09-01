package domain

import (
	"fmt"
	"time"
)

type Post struct {
	ID          string     `sql:"id,primary"`
	AuthorID    string     `sql:"author_id"`
	Title       string     `sql:"title"`
	Slug        string     `sql:"slug"`
	RawMarkdown string     `sql:"raw_markdown"`
	PublishedAt *time.Time `sql:"published_at"`
	CreatedAt   time.Time  `sql:"created_at"`
	UpdatedAt   time.Time  `sql:"updated_at"`
}

func NewPost(id string, authorID string, title string, slug string, rawMarkdown string) (*Post, error) {
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

	now := time.Now()
	return &Post{
		ID:          id,
		AuthorID:    authorID,
		Title:       title,
		Slug:        slug,
		RawMarkdown: rawMarkdown,
		PublishedAt: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func NewPublishedPost(id string, authorID string, title string, slug string, rawMarkdown string, publishedAt time.Time) (*Post, error) {
	post, err := NewPost(id, authorID, title, slug, rawMarkdown)
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

func (p *Post) Update(title string, slug string, rawMarkdown string) error {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	if slug == "" {
		return fmt.Errorf("slug cannot be empty")
	}

	if rawMarkdown == "" {
		return fmt.Errorf("raw markdown cannot be empty")
	}

	p.Title = title
	p.Slug = slug
	p.RawMarkdown = rawMarkdown
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Post) TableName() string {
	return "posts"
}
