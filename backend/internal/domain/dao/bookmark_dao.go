package dao

import (
	"blog0/internal/domain"
	"context"
)

type Bookmark = domain.Bookmark

type BookmarkDAO interface {
	// Create creates a new Bookmark
	Create(ctx context.Context, m *Bookmark) error

	// Update updates an existing Bookmark
	Update(ctx context.Context, m *Bookmark) error

	// PartialUpdate updates specific fields of a Bookmark
	PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error

	// DeleteByPk deletes a Bookmark by primary key
	DeleteByPk(ctx context.Context, pk string) error

	// FindByPk finds a Bookmark by primary key
	FindByPk(ctx context.Context, pk string) (*Bookmark, error)

	// CreateMany creates multiple Bookmark records
	CreateMany(ctx context.Context, models []*Bookmark) error

	// UpdateMany updates multiple Bookmark records
	UpdateMany(ctx context.Context, models []*Bookmark) error

	// DeleteManyByPks deletes multiple Bookmark records by primary keys
	DeleteManyByPks(ctx context.Context, pks []string) error

	// FindOne finds a single Bookmark with optional where clause
	FindOne(ctx context.Context, where string, args ...interface{}) (*Bookmark, error)

	// FindAll finds all Bookmark records with optional where clause
	FindAll(ctx context.Context, where string, args ...interface{}) ([]*Bookmark, error)

	// FindPaginated finds Bookmark records with pagination and optional where clause
	FindPaginated(ctx context.Context, limit, offset int, where string, args ...interface{}) ([]*Bookmark, error)

	// Count counts Bookmark records with optional where clause
	Count(ctx context.Context, where string, args ...interface{}) (int64, error)

	// WithTransaction executes a function within a database transaction
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
