package services

import (
	"context"
	"time"
)

type CreateComment struct{}

type CreateCommentReq struct {
	Slug     string  `json:"-"`
	Body     string  `json:"body"`
	ParentID *string `json:"parent_id"`
	UserID   string  `json:"-"`
}

type CreateCommentResp struct {
	ID       string     `json:"id"`
	PostSlug string     `json:"post_slug"`
	Author   AuthorInfo `json:"author"`
	ParentID *string    `json:"parent_id"`
	Body     string     `json:"body"`
	CreateAt time.Time  `json:"created_at"`
}

func NewCreateComment() *CreateComment {
	return &CreateComment{}
}

func (s *CreateComment) Exec(ctx context.Context, req *CreateCommentReq) (*CreateCommentResp, error) {
	return &CreateCommentResp{
		ID:       "uuid-nuevo-comentario",
		PostSlug: req.Slug,
		Author:   AuthorInfo{ID: req.UserID, Name: "Sample User"},
		ParentID: req.ParentID,
		Body:     req.Body,
		CreateAt: time.Now(),
	}, nil
}