package domain

import (
	"fmt"
	"time"
)

type Follow struct {
	FollowerID string    `sql:"follower_id"`
	FolloweeID string    `sql:"followee_id"`
	CreatedAt  time.Time `sql:"created_at"`
}

func NewFollow(followerID string, followeeID string) (*Follow, error) {
	if followerID == "" {
		return nil, fmt.Errorf("follower ID cannot be empty")
	}

	if followeeID == "" {
		return nil, fmt.Errorf("followee ID cannot be empty")
	}

	if followerID == followeeID {
		return nil, fmt.Errorf("cannot follow yourself")
	}

	return &Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
		CreatedAt:  time.Now(),
	}, nil
}

func (f *Follow) TableName() string {
	return "follows"
}
