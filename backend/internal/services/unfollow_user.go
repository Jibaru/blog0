package services

import (
	"context"
)

type UnfollowUser struct{}

type UnfollowUserReq struct {
	AuthorID string `json:"-"`
	UserID   string `json:"-"`
}

type UnfollowUserResp struct {
	Following      bool `json:"following"`
	FollowersCount int  `json:"followers_count"`
}

func NewUnfollowUser() *UnfollowUser {
	return &UnfollowUser{}
}

func (s *UnfollowUser) Exec(ctx context.Context, req *UnfollowUserReq) (*UnfollowUserResp, error) {
	return &UnfollowUserResp{
		Following:      false,
		FollowersCount: 127,
	}, nil
}