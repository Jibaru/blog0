package domain

import (
	"fmt"
	"time"
)

type Follow struct {
	ID         string    `sql:"id,primary"`
	FollowerID string    `sql:"follower_id"`
	FolloweeID string    `sql:"followee_id"`
	CreatedAt  time.Time `sql:"created_at"`
}

func NewFollow(id string, followerID string, followeeID string) (*Follow, error) {
	if id == "" {
		return nil, fmt.Errorf("ID cannot be empty")
	}

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
		ID:         id,
		FollowerID: followerID,
		FolloweeID: followeeID,
		CreatedAt:  time.Now(),
	}, nil
}

func (f *Follow) TableName() string {
	return "follows"
}
