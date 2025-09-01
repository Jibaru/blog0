package services

import (
	"context"
	"fmt"

	"blog0/internal/domain/dao"
)

type UnbookmarkPost struct {
	postDAO     dao.PostDAO
	bookmarkDAO dao.BookmarkDAO
}

type UnbookmarkPostReq struct {
	Slug   string `json:"-"`
	UserID string `json:"-"`
}

type UnbookmarkPostResp struct {
	Bookmarked bool `json:"bookmarked"`
}

func NewUnbookmarkPost(postDAO dao.PostDAO, bookmarkDAO dao.BookmarkDAO) *UnbookmarkPost {
	return &UnbookmarkPost{
		postDAO:     postDAO,
		bookmarkDAO: bookmarkDAO,
	}
}

func (s *UnbookmarkPost) Exec(ctx context.Context, req *UnbookmarkPostReq) (*UnbookmarkPostResp, error) {
	post, err := s.postDAO.FindOne(ctx, "slug = $1", "", req.Slug)
	if err != nil {
		return nil, fmt.Errorf("post not found: %w", err)
	}

	bookmark, err := s.bookmarkDAO.FindOne(ctx, "user_id = $1 AND post_id = $2", "", req.UserID, post.ID)
	if err != nil {
		return nil, fmt.Errorf("bookmark not found: %w", err)
	}

	err = s.bookmarkDAO.DeleteByPk(ctx, bookmark.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete bookmark: %w", err)
	}

	return &UnbookmarkPostResp{
		Bookmarked: false,
	}, nil
}