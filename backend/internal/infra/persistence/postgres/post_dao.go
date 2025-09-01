package postgres

import (
	"blog0/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Post = domain.Post

type PostDAO struct {
	db *sql.DB
}

func NewPostDAO(db *sql.DB) *PostDAO {
	return &PostDAO{db: db}
}

func (dao *PostDAO) getTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value("currentTx").(*sql.Tx); ok {
		return tx
	}
	return nil
}

func (dao *PostDAO) execContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return dao.db.ExecContext(ctx, query, args...)
}

func (dao *PostDAO) queryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return dao.db.QueryRowContext(ctx, query, args...)
}

func (dao *PostDAO) queryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return dao.db.QueryContext(ctx, query, args...)
}

func (dao *PostDAO) Create(ctx context.Context, m *Post) error {
	query := `
		INSERT INTO posts (id, author_id, title, slug, raw_markdown, published_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := dao.execContext(
		ctx,
		query,
		m.ID,
		m.AuthorID,
		m.Title,
		m.Slug,
		m.RawMarkdown,
		m.PublishedAt,
		m.CreatedAt,
		m.UpdatedAt,
	)

	return err
}

func (dao *PostDAO) Update(ctx context.Context, m *Post) error {
	query := `
		UPDATE posts
		SET author_id = $1,
			title = $2,
			slug = $3,
			raw_markdown = $4,
			published_at = $5,
			created_at = $6,
			updated_at = $7
		WHERE id = $8
	`

	_, err := dao.execContext(ctx, query,
		m.AuthorID,
		m.Title,
		m.Slug,
		m.RawMarkdown,
		m.PublishedAt,
		m.CreatedAt,
		m.UpdatedAt,
		m.ID,
	)
	return err
}

func (dao *PostDAO) PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error {
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

	query := fmt.Sprintf(`UPDATE posts SET %s WHERE id = $%d`, strings.Join(setClauses, ", "), i)

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *PostDAO) DeleteByPk(ctx context.Context, pk string) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := dao.execContext(ctx, query, pk)
	return err
}

func (dao *PostDAO) FindByPk(ctx context.Context, pk string) (*Post, error) {
	query := `
		SELECT id, author_id, title, slug, raw_markdown, published_at, created_at, updated_at
		FROM posts
		WHERE id = $1
	`
	row := dao.queryRowContext(ctx, query, pk)

	var m Post
	err := row.Scan(
		&m.ID,
		&m.AuthorID,
		&m.Title,
		&m.Slug,
		&m.RawMarkdown,
		&m.PublishedAt,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *PostDAO) CreateMany(ctx context.Context, models []*Post) error {
	if len(models) == 0 {
		return nil
	}

	placeholders := make([]string, len(models))
	args := make([]interface{}, 0, len(models)*8)

	for i, model := range models {
		placeholders[i] = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*8+1, i*8+2, i*8+3, i*8+4, i*8+5, i*8+6, i*8+7, i*8+8)

		args = append(args,
			model.ID,
			model.AuthorID,
			model.Title,
			model.Slug,
			model.RawMarkdown,
			model.PublishedAt,
			model.CreatedAt,
			model.UpdatedAt,
		)
	}

	query := fmt.Sprintf(`
		INSERT INTO posts (id, author_id, title, slug, raw_markdown, published_at, created_at, updated_at)
		VALUES %s
	`, strings.Join(placeholders, ", "))

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *PostDAO) UpdateMany(ctx context.Context, models []*Post) error {
	if len(models) == 0 {
		return nil
	}

	query := `
		UPDATE posts
		SET author_id = $1,
			title = $2,
			slug = $3,
			raw_markdown = $4,
			published_at = $5,
			created_at = $6,
			updated_at = $7
		WHERE id = $8
	`

	for _, model := range models {
		_, err := dao.execContext(ctx, query,
			model.AuthorID,
			model.Title,
			model.Slug,
			model.RawMarkdown,
			model.PublishedAt,
			model.CreatedAt,
			model.UpdatedAt,
			model.ID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao *PostDAO) DeleteManyByPks(ctx context.Context, pks []string) error {
	if len(pks) == 0 {
		return nil
	}

	placeholders := make([]string, len(pks))
	args := make([]interface{}, len(pks))
	for i, pk := range pks {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = pk
	}

	query := fmt.Sprintf(`DELETE FROM posts WHERE id IN (%s)`, strings.Join(placeholders, ","))
	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *PostDAO) FindOne(ctx context.Context, where string, sort string, args ...interface{}) (*Post, error) {
	query := `
		SELECT id, author_id, title, slug, raw_markdown, published_at, created_at, updated_at
		FROM posts
	`

	if where != "" {
		query += " WHERE " + where
	}

	if sort != "" {
		query += " ORDER BY " + sort
	}

	row := dao.queryRowContext(ctx, query, args...)

	var m Post
	err := row.Scan(
		&m.ID,
		&m.AuthorID,
		&m.Title,
		&m.Slug,
		&m.RawMarkdown,
		&m.PublishedAt,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *PostDAO) FindAll(ctx context.Context, where string, sort string, args ...interface{}) ([]*Post, error) {
	query := `
		SELECT id, author_id, title, slug, raw_markdown, published_at, created_at, updated_at
		FROM posts
	`

	if where != "" {
		query += " WHERE " + where
	}

	if sort != "" {
		query += " ORDER BY " + sort
	}

	rows, err := dao.queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*Post
	for rows.Next() {
		var m Post
		err := rows.Scan(
			&m.ID,
			&m.AuthorID,
			&m.Title,
			&m.Slug,
			&m.RawMarkdown,
			&m.PublishedAt,
			&m.CreatedAt,
			&m.UpdatedAt,
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

func (dao *PostDAO) FindPaginated(ctx context.Context, limit, offset int, where string, sort string, args ...interface{}) ([]*Post, error) {
	query := `
		SELECT id, author_id, title, slug, raw_markdown, published_at, created_at, updated_at
		FROM posts
	`

	if where != "" {
		query += " WHERE " + where
	}

	if sort != "" {
		query += " ORDER BY " + sort
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := dao.queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*Post
	for rows.Next() {
		var m Post
		err := rows.Scan(
			&m.ID,
			&m.AuthorID,
			&m.Title,
			&m.Slug,
			&m.RawMarkdown,
			&m.PublishedAt,
			&m.CreatedAt,
			&m.UpdatedAt,
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

func (dao *PostDAO) Count(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := "SELECT COUNT(*) FROM posts"

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

func (dao *PostDAO) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
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
