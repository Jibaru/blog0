package services

import (
	"context"
	"fmt"

	"blog0/internal/domain/dao"
)

type DeletePost struct {
	postDAO dao.PostDAO
}

type DeletePostReq struct {
	Slug   string `json:"-"`
	UserID string `json:"-"`
}

type DeletePostResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewDeletePost(postDAO dao.PostDAO) *DeletePost {
	return &DeletePost{
		postDAO: postDAO,
	}
}

func (s *DeletePost) Exec(ctx context.Context, req *DeletePostReq) (*DeletePostResp, error) {
	post, err := s.postDAO.FindOne(ctx, "slug = $1", "", req.Slug)
	if err != nil {
		return nil, fmt.Errorf("post not found: %w", err)
	}

	if post.AuthorID != req.UserID {
		return nil, fmt.Errorf("unauthorized: you can only delete your own posts")
	}

	err = s.postDAO.DeleteByPk(ctx, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	return &DeletePostResp{
		Success: true,
		Message: "Post deleted successfully",
	}, nil
}