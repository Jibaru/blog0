package dao

import (
	"blog0/internal/domain"
	"context"
)

type PostLike = domain.PostLike

type PostLikeDAO interface {
	// Create creates a new PostLike
	Create(ctx context.Context, m *PostLike) error

	// Update updates an existing PostLike
	Update(ctx context.Context, m *PostLike) error

	// PartialUpdate updates specific fields of a PostLike
	PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error

	// DeleteByPk deletes a PostLike by primary key
	DeleteByPk(ctx context.Context, pk string) error

	// FindByPk finds a PostLike by primary key
	FindByPk(ctx context.Context, pk string) (*PostLike, error)

	// CreateMany creates multiple PostLike records
	CreateMany(ctx context.Context, models []*PostLike) error

	// UpdateMany updates multiple PostLike records
	UpdateMany(ctx context.Context, models []*PostLike) error

	// DeleteManyByPks deletes multiple PostLike records by primary keys
	DeleteManyByPks(ctx context.Context, pks []string) error

	// FindOne finds a single PostLike with optional where clause
	FindOne(ctx context.Context, where string, args ...interface{}) (*PostLike, error)

	// FindAll finds all PostLike records with optional where clause
	FindAll(ctx context.Context, where string, args ...interface{}) ([]*PostLike, error)

	// FindPaginated finds PostLike records with pagination and optional where clause
	FindPaginated(ctx context.Context, limit, offset int, where string, args ...interface{}) ([]*PostLike, error)

	// Count counts PostLike records with optional where clause
	Count(ctx context.Context, where string, args ...interface{}) (int64, error)

	// WithTransaction executes a function within a database transaction
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
