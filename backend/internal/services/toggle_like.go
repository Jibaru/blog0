package services

import (
	"context"
	"fmt"

	"blog0/internal/domain"
	"blog0/internal/domain/dao"
)

type ToggleLike struct {
	postDAO     dao.PostDAO
	postLikeDAO dao.PostLikeDAO
	nextID      domain.NextID
}

type ToggleLikeReq struct {
	Slug   string `json:"-"`
	UserID string `json:"-"`
}

type ToggleLikeResp struct {
	Liked      bool `json:"liked"`
	LikesCount int  `json:"likes_count"`
}

func NewToggleLike(postDAO dao.PostDAO, postLikeDAO dao.PostLikeDAO, nextID domain.NextID) *ToggleLike {
	return &ToggleLike{
		postDAO:     postDAO,
		postLikeDAO: postLikeDAO,
		nextID:      nextID,
	}
}

func (s *ToggleLike) Exec(ctx context.Context, req *ToggleLikeReq) (*ToggleLikeResp, error) {
	post, err := s.postDAO.FindOne(ctx, "slug = $1", "", req.Slug)
	if err != nil {
		return nil, fmt.Errorf("post not found: %w", err)
	}

	existingLike, err := s.postLikeDAO.FindOne(ctx, "user_id = $1 AND post_id = $2", "", req.UserID, post.ID)
	
	var liked bool
	if err == nil && existingLike != nil {
		err = s.postLikeDAO.DeleteByPk(ctx, existingLike.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to remove like: %w", err)
		}
		liked = false
	} else {
		likeID := s.nextID()
		newLike, err := domain.NewPostLike(likeID, req.UserID, post.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to create like: %w", err)
		}
		
		err = s.postLikeDAO.Create(ctx, newLike)
		if err != nil {
			return nil, fmt.Errorf("failed to save like: %w", err)
		}
		liked = true
	}

	likesCount, err := s.postLikeDAO.Count(ctx, "post_id = $1", post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to count likes: %w", err)
	}

	return &ToggleLikeResp{
		Liked:      liked,
		LikesCount: int(likesCount),
	}, nil
}