package dao

import (
	"blog0/internal/domain"
	"context"
)

type Follow = domain.Follow

type FollowDAO interface {
	// Create creates a new Follow
	Create(ctx context.Context, m *Follow) error

	// Update updates an existing Follow
	Update(ctx context.Context, m *Follow) error

	// PartialUpdate updates specific fields of a Follow
	PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error

	// DeleteByPk deletes a Follow by primary key
	DeleteByPk(ctx context.Context, pk string) error

	// FindByPk finds a Follow by primary key
	FindByPk(ctx context.Context, pk string) (*Follow, error)

	// CreateMany creates multiple Follow records
	CreateMany(ctx context.Context, models []*Follow) error

	// UpdateMany updates multiple Follow records
	UpdateMany(ctx context.Context, models []*Follow) error

	// DeleteManyByPks deletes multiple Follow records by primary keys
	DeleteManyByPks(ctx context.Context, pks []string) error

	// FindOne finds a single Follow with optional where clause and sort expression
	FindOne(ctx context.Context, where string, sort string, args ...interface{}) (*Follow, error)

	// FindAll finds all Follow records with optional where clause and sort expression
	FindAll(ctx context.Context, where string, sort string, args ...interface{}) ([]*Follow, error)

	// FindPaginated finds Follow records with pagination, optional where clause and sort expression
	FindPaginated(ctx context.Context, limit, offset int, where string, sort string, args ...interface{}) ([]*Follow, error)

	// Count counts Follow records with optional where clause
	Count(ctx context.Context, where string, args ...interface{}) (int64, error)

	// WithTransaction executes a function within a database transaction
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
