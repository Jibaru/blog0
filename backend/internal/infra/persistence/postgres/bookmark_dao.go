package postgres

import (
	"blog0/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Bookmark = domain.Bookmark

type BookmarkDAO struct {
	db *sql.DB
}

func NewBookmarkDAO(db *sql.DB) *BookmarkDAO {
	return &BookmarkDAO{db: db}
}

func (dao *BookmarkDAO) getTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value("currentTx").(*sql.Tx); ok {
		return tx
	}
	return nil
}

func (dao *BookmarkDAO) execContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return dao.db.ExecContext(ctx, query, args...)
}

func (dao *BookmarkDAO) queryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return dao.db.QueryRowContext(ctx, query, args...)
}

func (dao *BookmarkDAO) queryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return dao.db.QueryContext(ctx, query, args...)
}

func (dao *BookmarkDAO) Create(ctx context.Context, m *Bookmark) error {
	query := `
		INSERT INTO bookmarks (id, user_id, post_id, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := dao.execContext(
		ctx,
		query,
		m.ID,
		m.UserID,
		m.PostID,
		m.CreatedAt,
	)

	return err
}

func (dao *BookmarkDAO) Update(ctx context.Context, m *Bookmark) error {
	query := `
		UPDATE bookmarks
		SET user_id = $1,
			post_id = $2,
			created_at = $3
		WHERE id = $4
	`

	_, err := dao.execContext(ctx, query,
		m.UserID,
		m.PostID,
		m.CreatedAt,
		m.ID,
	)
	return err
}

func (dao *BookmarkDAO) PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	setClauses := make([]string, 0, len(fields))
	args := make([]interface{}, 0, len(fields)+1)
	i := 1

	for field, value := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}

	args = append(args, pk)

	query := fmt.Sprintf(`UPDATE bookmarks SET %s WHERE id = $%d`, strings.Join(setClauses, ", "), i)

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *BookmarkDAO) DeleteByPk(ctx context.Context, pk string) error {
	query := `DELETE FROM bookmarks WHERE id = $1`
	_, err := dao.execContext(ctx, query, pk)
	return err
}

func (dao *BookmarkDAO) FindByPk(ctx context.Context, pk string) (*Bookmark, error) {
	query := `
		SELECT id, user_id, post_id, created_at
		FROM bookmarks
		WHERE id = $1
	`
	row := dao.queryRowContext(ctx, query, pk)

	var m Bookmark
	err := row.Scan(
		&m.ID,
		&m.UserID,
		&m.PostID,
		&m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *BookmarkDAO) CreateMany(ctx context.Context, models []*Bookmark) error {
	if len(models) == 0 {
		return nil
	}

	placeholders := make([]string, len(models))
	args := make([]interface{}, 0, len(models)*4)

	for i, model := range models {
		placeholders[i] = fmt.Sprintf("($%d, $%d, $%d, $%d)",
			i*4+1, i*4+2, i*4+3, i*4+4)

		args = append(args,
			model.ID,
			model.UserID,
			model.PostID,
			model.CreatedAt,
		)
	}

	query := fmt.Sprintf(`
		INSERT INTO bookmarks (id, user_id, post_id, created_at)
		VALUES %s
	`, strings.Join(placeholders, ", "))

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *BookmarkDAO) UpdateMany(ctx context.Context, models []*Bookmark) error {
	if len(models) == 0 {
		return nil
	}

	query := `
		UPDATE bookmarks
		SET user_id = $1,
			post_id = $2,
			created_at = $3
		WHERE id = $4
	`

	for _, model := range models {
		_, err := dao.execContext(ctx, query,
			model.UserID,
			model.PostID,
			model.CreatedAt,
			model.ID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao *BookmarkDAO) DeleteManyByPks(ctx context.Context, pks []string) error {
	if len(pks) == 0 {
		return nil
	}

	placeholders := make([]string, len(pks))
	args := make([]interface{}, len(pks))
	for i, pk := range pks {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = pk
	}

	query := fmt.Sprintf(`DELETE FROM bookmarks WHERE id IN (%s)`, strings.Join(placeholders, ","))
	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *BookmarkDAO) FindOne(ctx context.Context, where string, args ...interface{}) (*Bookmark, error) {
	query := `
		SELECT id, user_id, post_id, created_at
		FROM bookmarks
	`

	if where != "" {
		query += " WHERE " + where
	}

	row := dao.queryRowContext(ctx, query, args...)

	var m Bookmark
	err := row.Scan(
		&m.ID,
		&m.UserID,
		&m.PostID,
		&m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *BookmarkDAO) FindAll(ctx context.Context, where string, args ...interface{}) ([]*Bookmark, error) {
	query := `
		SELECT id, user_id, post_id, created_at
		FROM bookmarks
	`

	if where != "" {
		query += " WHERE " + where
	}

	rows, err := dao.queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*Bookmark
	for rows.Next() {
		var m Bookmark
		err := rows.Scan(
			&m.ID,
			&m.UserID,
			&m.PostID,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (dao *BookmarkDAO) FindPaginated(ctx context.Context, limit, offset int, where string, args ...interface{}) ([]*Bookmark, error) {
	query := `
		SELECT id, user_id, post_id, created_at
		FROM bookmarks
	`

	if where != "" {
		query += " WHERE " + where
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := dao.queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*Bookmark
	for rows.Next() {
		var m Bookmark
		err := rows.Scan(
			&m.ID,
			&m.UserID,
			&m.PostID,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (dao *BookmarkDAO) Count(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := "SELECT COUNT(*) FROM bookmarks"

	if where != "" {
		query += " WHERE " + where
	}

	row := dao.queryRowContext(ctx, query, args...)

	var count int64
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *BookmarkDAO) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ctxWithTx := context.WithValue(ctx, "currentTx", tx)

	err = fn(ctxWithTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
