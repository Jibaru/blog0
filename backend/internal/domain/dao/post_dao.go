package dao

import (
	"blog0/internal/domain"
	"context"
)

type Post = domain.Post

type PostDAO interface {
	// Create creates a new Post
	Create(ctx context.Context, m *Post) error

	// Update updates an existing Post
	Update(ctx context.Context, m *Post) error

	// PartialUpdate updates specific fields of a Post
	PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error

	// DeleteByPk deletes a Post by primary key
	DeleteByPk(ctx context.Context, pk string) error

	// FindByPk finds a Post by primary key
	FindByPk(ctx context.Context, pk string) (*Post, error)

	// CreateMany creates multiple Post records
	CreateMany(ctx context.Context, models []*Post) error

	// UpdateMany updates multiple Post records
	UpdateMany(ctx context.Context, models []*Post) error

	// DeleteManyByPks deletes multiple Post records by primary keys
	DeleteManyByPks(ctx context.Context, pks []string) error

	// FindOne finds a single Post with optional where clause and sort expression
	FindOne(ctx context.Context, where string, sort string, args ...interface{}) (*Post, error)

	// FindAll finds all Post records with optional where clause and sort expression
	FindAll(ctx context.Context, where string, sort string, args ...interface{}) ([]*Post, error)

	// FindPaginated finds Post records with pagination, optional where clause and sort expression
	FindPaginated(ctx context.Context, limit, offset int, where string, sort string, args ...interface{}) ([]*Post, error)

	// Count counts Post records with optional where clause
	Count(ctx context.Context, where string, args ...interface{}) (int64, error)

	// WithTransaction executes a function within a database transaction
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
