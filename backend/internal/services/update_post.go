package services

import (
	"context"
	"fmt"
	"time"

	"blog0/internal/domain/dao"
)

type UpdatePost struct {
	postDAO dao.PostDAO
}

type UpdatePostReq struct {
	Slug        string `json:"-"`
	Title       string `json:"title"`
	NewSlug     string `json:"slug"`
	RawMarkdown string `json:"raw_markdown"`
	UserID      string `json:"-"`
	Publish     *bool  `json:"publish"`
}

type UpdatePostResp struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	RawMarkdown string     `json:"raw_markdown"`
	AuthorID    string     `json:"author_id"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewUpdatePost(postDAO dao.PostDAO) *UpdatePost {
	return &UpdatePost{
		postDAO: postDAO,
	}
}

func (s *UpdatePost) Exec(ctx context.Context, req *UpdatePostReq) (*UpdatePostResp, error) {
	post, err := s.postDAO.FindOne(ctx, "slug = $1", "", req.Slug)
	if err != nil {
		return nil, fmt.Errorf("post not found: %w", err)
	}

	if post.AuthorID != req.UserID {
		return nil, fmt.Errorf("unauthorized: you can only update your own posts")
	}

	if req.Title != "" && req.NewSlug != "" && req.RawMarkdown != "" {
		err = post.Update(req.Title, req.NewSlug, req.RawMarkdown)
		if err != nil {
			return nil, fmt.Errorf("failed to update post: %w", err)
		}
	}

	if req.Publish != nil {
		if *req.Publish && post.PublishedAt == nil {
			post.Publish(time.Now())
		} else if !*req.Publish && post.PublishedAt != nil {
			post.PublishedAt = nil
			post.UpdatedAt = time.Now()
		}
	}

	err = s.postDAO.Update(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("failed to save post: %w", err)
	}

	return &UpdatePostResp{
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