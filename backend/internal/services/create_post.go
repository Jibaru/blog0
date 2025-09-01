package services

import (
	"context"
	"fmt"
	"time"

	"blog0/internal/domain"
	"blog0/internal/domain/dao"
)

type CreatePost struct {
	postDAO dao.PostDAO
	nextID  domain.NextID
}

type CreatePostReq struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	RawMarkdown string `json:"raw_markdown"`
	UserID      string `json:"-"`
	Publish     bool   `json:"publish"`
}

type CreatePostResp struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	RawMarkdown string     `json:"raw_markdown"`
	AuthorID    string     `json:"author_id"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewCreatePost(postDAO dao.PostDAO, nextID domain.NextID) *CreatePost {
	return &CreatePost{
		postDAO: postDAO,
		nextID:  nextID,
	}
}

func (s *CreatePost) Exec(ctx context.Context, req *CreatePostReq) (*CreatePostResp, error) {
	postID := s.nextID()

	var post *domain.Post
	var err error

	if req.Publish {
		publishedAt := time.Now()
		post, err = domain.NewPublishedPost(postID, req.UserID, req.Title, req.Slug, req.RawMarkdown, publishedAt)
	} else {
		post, err = domain.NewPost(postID, req.UserID, req.Title, req.Slug, req.RawMarkdown)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	err = s.postDAO.Create(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("failed to save post: %w", err)
	}

	return &CreatePostResp{
		ID:          post.ID,
		Title:       post.Title,
		Slug:        post.Slug,
		RawMarkdown: post.RawMarkdown,
		AuthorID:    post.AuthorID,
		PublishedAt: post.PublishedAt,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}, nil
}