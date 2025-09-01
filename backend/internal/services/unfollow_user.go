package services

import (
	"context"
	"fmt"

	"blog0/internal/domain/dao"
)

type UnfollowUser struct {
	userDAO   dao.UserDAO
	followDAO dao.FollowDAO
}

type UnfollowUserReq struct {
	AuthorID string `json:"-"`
	UserID   string `json:"-"`
}

type UnfollowUserResp struct {
	Following      bool `json:"following"`
	FollowersCount int  `json:"followers_count"`
}

func NewUnfollowUser(userDAO dao.UserDAO, followDAO dao.FollowDAO) *UnfollowUser {
	return &UnfollowUser{
		userDAO:   userDAO,
		followDAO: followDAO,
	}
}

func (s *UnfollowUser) Exec(ctx context.Context, req *UnfollowUserReq) (*UnfollowUserResp, error) {
	_, err := s.userDAO.FindByPk(ctx, req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	follow, err := s.followDAO.FindByComposite(ctx, req.UserID, req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("follow relationship not found: %w", err)
	}

	err = s.followDAO.DeleteByComposite(ctx, follow.FollowerID, follow.FolloweeID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete follow: %w", err)
	}

	followersCount, err := s.followDAO.Count(ctx, "followee_id = $1", req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to count followers: %w", err)
	}

	return &UnfollowUserResp{
		Following:      false,
		FollowersCount: int(followersCount),
	}, nil
}