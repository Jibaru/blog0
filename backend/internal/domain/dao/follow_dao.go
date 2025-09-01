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

	// DeleteByComposite deletes a Follow by composite primary key (follower_id, followee_id)
	DeleteByComposite(ctx context.Context, followerID, followeeID string) error

	// FindByComposite finds a Follow by composite primary key
	FindByComposite(ctx context.Context, followerID, followeeID string) (*Follow, error)

	// CreateMany creates multiple Follow records
	CreateMany(ctx context.Context, models []*Follow) error

	// UpdateMany updates multiple Follow records
	UpdateMany(ctx context.Context, models []*Follow) error

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