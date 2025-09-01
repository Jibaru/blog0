package domain

import (
	"fmt"
	"time"
)

type Comment struct {
	ID        string    `sql:"id,primary"`
	PostID    string    `sql:"post_id"`
	AuthorID  string    `sql:"author_id"`
	ParentID  *string   `sql:"parent_id"`
	Body      string    `sql:"body"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
}

func NewComment(id string, postID string, authorID string, body string) (*Comment, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	if postID == "" {
		return nil, fmt.Errorf("post ID cannot be empty")
	}

	if authorID == "" {
		return nil, fmt.Errorf("author ID cannot be empty")
	}

	if body == "" {
		return nil, fmt.Errorf("body cannot be empty")
	}

	now := time.Now()
	return &Comment{
		ID:        id,
		PostID:    postID,
		AuthorID:  authorID,
		ParentID:  nil,
		Body:      body,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func NewReplyComment(id string, postID string, authorID string, parentID string, body string) (*Comment, error) {
	comment, err := NewComment(id, postID, authorID, body)
	if err != nil {
		return nil, err
	}

	if parentID == "" {
		return nil, fmt.Errorf("parent ID cannot be empty for reply comment")
	}

	comment.ParentID = &parentID
	return comment, nil
}

func (c *Comment) UpdateBody(body string) error {
	if body == "" {
		return fmt.Errorf("body cannot be empty")
	}

	c.Body = body
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Comment) TableName() string {
	return "comments"
}
