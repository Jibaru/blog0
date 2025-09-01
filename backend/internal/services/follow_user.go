package services

import (
	"context"
	"fmt"

	"blog0/internal/domain"
	"blog0/internal/domain/dao"
)

type FollowUser struct {
	userDAO   dao.UserDAO
	followDAO dao.FollowDAO
}

type FollowUserReq struct {
	AuthorID string `json:"-"`
	UserID   string `json:"-"`
}

type FollowUserResp struct {
	Following      bool `json:"following"`
	FollowersCount int  `json:"followers_count"`
}

func NewFollowUser(userDAO dao.UserDAO, followDAO dao.FollowDAO) *FollowUser {
	return &FollowUser{
		userDAO:   userDAO,
		followDAO: followDAO,
	}
}

func (s *FollowUser) Exec(ctx context.Context, req *FollowUserReq) (*FollowUserResp, error) {
	_, err := s.userDAO.FindByPk(ctx, req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	_, err = s.userDAO.FindByPk(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	existingFollow, err := s.followDAO.FindByComposite(ctx, req.UserID, req.AuthorID)
	if err == nil && existingFollow != nil {
		return nil, fmt.Errorf("user is already following this author")
	}

	newFollow, err := domain.NewFollow(req.UserID, req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to create follow: %w", err)
	}

	err = s.followDAO.Create(ctx, newFollow)
	if err != nil {
		return nil, fmt.Errorf("failed to save follow: %w", err)
	}

	followersCount, err := s.followDAO.Count(ctx, "followee_id = $1", req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to count followers: %w", err)
	}

	return &FollowUserResp{
		Following:      true,
		FollowersCount: int(followersCount),
	}, nil
}