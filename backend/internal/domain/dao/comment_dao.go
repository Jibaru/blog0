package dao

import (
	"blog0/internal/domain"
	"context"
)

type Comment = domain.Comment

type CommentDAO interface {
	// Create creates a new Comment
	Create(ctx context.Context, m *Comment) error

	// Update updates an existing Comment
	Update(ctx context.Context, m *Comment) error

	// PartialUpdate updates specific fields of a Comment
	PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error

	// DeleteByPk deletes a Comment by primary key
	DeleteByPk(ctx context.Context, pk string) error

	// FindByPk finds a Comment by primary key
	FindByPk(ctx context.Context, pk string) (*Comment, error)

	// CreateMany creates multiple Comment records
	CreateMany(ctx context.Context, models []*Comment) error

	// UpdateMany updates multiple Comment records
	UpdateMany(ctx context.Context, models []*Comment) error

	// DeleteManyByPks deletes multiple Comment records by primary keys
	DeleteManyByPks(ctx context.Context, pks []string) error

	// FindOne finds a single Comment with optional where clause and sort expression
	FindOne(ctx context.Context, where string, sort string, args ...interface{}) (*Comment, error)

	// FindAll finds all Comment records with optional where clause and sort expression
	FindAll(ctx context.Context, where string, sort string, args ...interface{}) ([]*Comment, error)

	// FindPaginated finds Comment records with pagination, optional where clause and sort expression
	FindPaginated(ctx context.Context, limit, offset int, where string, sort string, args ...interface{}) ([]*Comment, error)

	// Count counts Comment records with optional where clause
	Count(ctx context.Context, where string, args ...interface{}) (int64, error)

	// WithTransaction executes a function within a database transaction
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
