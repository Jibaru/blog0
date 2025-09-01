package postgres

import (
	"blog0/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Comment = domain.Comment

type CommentDAO struct {
	db *sql.DB
}

func NewCommentDAO(db *sql.DB) *CommentDAO {
	return &CommentDAO{db: db}
}

func (dao *CommentDAO) getTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value("currentTx").(*sql.Tx); ok {
		return tx
	}
	return nil
}

func (dao *CommentDAO) execContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return dao.db.ExecContext(ctx, query, args...)
}

func (dao *CommentDAO) queryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return dao.db.QueryRowContext(ctx, query, args...)
}

func (dao *CommentDAO) queryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return dao.db.QueryContext(ctx, query, args...)
}

func (dao *CommentDAO) Create(ctx context.Context, m *Comment) error {
	query := `
		INSERT INTO comments (id, post_id, author_id, parent_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := dao.execContext(
		ctx,
		query,
		m.ID,
		m.PostID,
		m.AuthorID,
		m.ParentID,
		m.Body,
		m.CreatedAt,
		m.UpdatedAt,
	)

	return err
}

func (dao *CommentDAO) Update(ctx context.Context, m *Comment) error {
	query := `
		UPDATE comments
		SET post_id = $1,
			author_id = $2,
			parent_id = $3,
			body = $4,
			created_at = $5,
			updated_at = $6
		WHERE id = $7
	`

	_, err := dao.execContext(ctx, query,
		m.PostID,
		m.AuthorID,
		m.ParentID,
		m.Body,
		m.CreatedAt,
		m.UpdatedAt,
		m.ID,
	)
	return err
}

func (dao *CommentDAO) PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error {
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

	query := fmt.Sprintf(`UPDATE comments SET %s WHERE id = $%d`, strings.Join(setClauses, ", "), i)

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *CommentDAO) DeleteByPk(ctx context.Context, pk string) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := dao.execContext(ctx, query, pk)
	return err
}

func (dao *CommentDAO) FindByPk(ctx context.Context, pk string) (*Comment, error) {
	query := `
		SELECT id, post_id, author_id, parent_id, body, created_at, updated_at
		FROM comments
		WHERE id = $1
	`
	row := dao.queryRowContext(ctx, query, pk)

	var m Comment
	err := row.Scan(
		&m.ID,
		&m.PostID,
		&m.AuthorID,
		&m.ParentID,
		&m.Body,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *CommentDAO) CreateMany(ctx context.Context, models []*Comment) error {
	if len(models) == 0 {
		return nil
	}

	placeholders := make([]string, len(models))
	args := make([]interface{}, 0, len(models)*7)

	for i, model := range models {
		placeholders[i] = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*7+1, i*7+2, i*7+3, i*7+4, i*7+5, i*7+6, i*7+7)

		args = append(args,
			model.ID,
			model.PostID,
			model.AuthorID,
			model.ParentID,
			model.Body,
			model.CreatedAt,
			model.UpdatedAt,
		)
	}

	query := fmt.Sprintf(`
		INSERT INTO comments (id, post_id, author_id, parent_id, body, created_at, updated_at)
		VALUES %s
	`, strings.Join(placeholders, ", "))

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *CommentDAO) UpdateMany(ctx context.Context, models []*Comment) error {
	if len(models) == 0 {
		return nil
	}

	query := `
		UPDATE comments
		SET post_id = $1,
			author_id = $2,
			parent_id = $3,
			body = $4,
			created_at = $5,
			updated_at = $6
		WHERE id = $7
	`

	for _, model := range models {
		_, err := dao.execContext(ctx, query,
			model.PostID,
			model.AuthorID,
			model.ParentID,
			model.Body,
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

func (dao *CommentDAO) DeleteManyByPks(ctx context.Context, pks []string) error {
	if len(pks) == 0 {
		return nil
	}

	placeholders := make([]string, len(pks))
	args := make([]interface{}, len(pks))
	for i, pk := range pks {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = pk
	}

	query := fmt.Sprintf(`DELETE FROM comments WHERE id IN (%s)`, strings.Join(placeholders, ","))
	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *CommentDAO) FindOne(ctx context.Context, where string, args ...interface{}) (*Comment, error) {
	query := `
		SELECT id, post_id, author_id, parent_id, body, created_at, updated_at
		FROM comments
	`

	if where != "" {
		query += " WHERE " + where
	}

	row := dao.queryRowContext(ctx, query, args...)

	var m Comment
	err := row.Scan(
		&m.ID,
		&m.PostID,
		&m.AuthorID,
		&m.ParentID,
		&m.Body,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *CommentDAO) FindAll(ctx context.Context, where string, args ...interface{}) ([]*Comment, error) {
	query := `
		SELECT id, post_id, author_id, parent_id, body, created_at, updated_at
		FROM comments
	`

	if where != "" {
		query += " WHERE " + where
	}

	rows, err := dao.queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*Comment
	for rows.Next() {
		var m Comment
		err := rows.Scan(
			&m.ID,
			&m.PostID,
			&m.AuthorID,
			&m.ParentID,
			&m.Body,
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

func (dao *CommentDAO) FindPaginated(ctx context.Context, limit, offset int, where string, args ...interface{}) ([]*Comment, error) {
	query := `
		SELECT id, post_id, author_id, parent_id, body, created_at, updated_at
		FROM comments
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

	var models []*Comment
	for rows.Next() {
		var m Comment
		err := rows.Scan(
			&m.ID,
			&m.PostID,
			&m.AuthorID,
			&m.ParentID,
			&m.Body,
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

func (dao *CommentDAO) Count(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := "SELECT COUNT(*) FROM comments"

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

func (dao *CommentDAO) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
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
