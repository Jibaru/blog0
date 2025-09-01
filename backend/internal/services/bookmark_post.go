package services

import (
	"context"
	"fmt"
	"time"

	"blog0/internal/domain"
	"blog0/internal/domain/dao"
)

type BookmarkPost struct {
	postDAO     dao.PostDAO
	bookmarkDAO dao.BookmarkDAO
	nextID      domain.NextID
}

type BookmarkPostReq struct {
	Slug   string `json:"-"`
	UserID string `json:"-"`
}

type BookmarkPostResp struct {
	Bookmarked bool      `json:"bookmarked"`
	BookmarkID string    `json:"bookmark_id"`
	PostSlug   string    `json:"post_slug"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewBookmarkPost(postDAO dao.PostDAO, bookmarkDAO dao.BookmarkDAO, nextID domain.NextID) *BookmarkPost {
	return &BookmarkPost{
		postDAO:     postDAO,
		bookmarkDAO: bookmarkDAO,
		nextID:      nextID,
	}
}

func (s *BookmarkPost) Exec(ctx context.Context, req *BookmarkPostReq) (*BookmarkPostResp, error) {
	post, err := s.postDAO.FindOne(ctx, "slug = $1", "", req.Slug)
	if err != nil {
		return nil, fmt.Errorf("post not found: %w", err)
	}

	existingBookmark, err := s.bookmarkDAO.FindOne(ctx, "user_id = $1 AND post_id = $2", "", req.UserID, post.ID)
	if err == nil && existingBookmark != nil {
		return nil, fmt.Errorf("post is already bookmarked")
	}

	bookmarkID := s.nextID()
	newBookmark, err := domain.NewBookmark(bookmarkID, req.UserID, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create bookmark: %w", err)
	}

	err = s.bookmarkDAO.Create(ctx, newBookmark)
	if err != nil {
		return nil, fmt.Errorf("failed to save bookmark: %w", err)
	}

	return &BookmarkPostResp{
		Bookmarked: true,
		BookmarkID: newBookmark.ID,
		PostSlug:   req.Slug,
		CreatedAt:  newBookmark.CreatedAt,
	}, nil
}