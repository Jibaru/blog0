package domain

import (
	"fmt"
	"time"
)

type Bookmark struct {
	ID        string    `sql:"id,primary"`
	UserID    string    `sql:"user_id"`
	PostID    string    `sql:"post_id"`
	CreatedAt time.Time `sql:"created_at"`
}

func NewBookmark(id string, userID string, postID string) (*Bookmark, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	if userID == "" {
		return nil, fmt.Errorf("user ID cannot be empty")
	}

	if postID == "" {
		return nil, fmt.Errorf("post ID cannot be empty")
	}

	return &Bookmark{
		ID:        id,
		UserID:    userID,
		PostID:    postID,
		CreatedAt: time.Now(),
	}, nil
}

func (b *Bookmark) TableName() string {
	return "bookmarks"
}
