package domain

import (
	"fmt"
	"time"
)

type PostLike struct {
	ID        string    `sql:"id,primary"`
	UserID    string    `sql:"user_id"`
	PostID    string    `sql:"post_id"`
	CreatedAt time.Time `sql:"created_at"`
}

func NewPostLike(id string, userID string, postID string) (*PostLike, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	if userID == "" {
		return nil, fmt.Errorf("user ID cannot be empty")
	}

	if postID == "" {
		return nil, fmt.Errorf("post ID cannot be empty")
	}

	return &PostLike{
		ID:        id,
		UserID:    userID,
		PostID:    postID,
		CreatedAt: time.Now(),
	}, nil
}

func (p *PostLike) TableName() string {
	return "post_likes"
}
