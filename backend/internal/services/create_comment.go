package services

import (
	"context"
	"fmt"
	"time"

	"blog0/internal/domain"
	"blog0/internal/domain/dao"
)

type CreateComment struct {
	postDAO    dao.PostDAO
	userDAO    dao.UserDAO
	commentDAO dao.CommentDAO
	nextID     domain.NextID
}

type CreateCommentReq struct {
	Slug     string  `json:"-"`
	Body     string  `json:"body"`
	ParentID *string `json:"parent_id"`
	UserID   string  `json:"-"`
}

type CreateCommentResp struct {
	ID        string     `json:"id"`
	PostSlug  string     `json:"post_slug"`
	Author    AuthorInfo `json:"author"`
	ParentID  *string    `json:"parent_id"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
}

func NewCreateComment(postDAO dao.PostDAO, userDAO dao.UserDAO, commentDAO dao.CommentDAO, nextID domain.NextID) *CreateComment {
	return &CreateComment{
		postDAO:    postDAO,
		userDAO:    userDAO,
		commentDAO: commentDAO,
		nextID:     nextID,
	}
}

func (s *CreateComment) Exec(ctx context.Context, req *CreateCommentReq) (*CreateCommentResp, error) {
	post, err := s.postDAO.FindOne(ctx, "slug = $1", "", req.Slug)
	if err != nil {
		return nil, fmt.Errorf("post not found: %w", err)
	}

	author, err := s.userDAO.FindByPk(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	commentID := s.nextID()

	var comment *domain.Comment
	if req.ParentID != nil {
		comment, err = domain.NewReplyComment(commentID, post.ID, req.UserID, *req.ParentID, req.Body)
	} else {
		comment, err = domain.NewComment(commentID, post.ID, req.UserID, req.Body)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}

	err = s.commentDAO.Create(ctx, comment)
	if err != nil {
		return nil, fmt.Errorf("failed to save comment: %w", err)
	}

	return &CreateCommentResp{
		ID:        comment.ID,
		PostSlug:  req.Slug,
		Author:    AuthorInfo{ID: author.ID, Name: author.Username},
		ParentID:  comment.ParentID,
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt,
	}, nil
}
