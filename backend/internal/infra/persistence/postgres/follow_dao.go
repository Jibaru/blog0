package postgres

import (
	"blog0/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Follow = domain.Follow

type FollowDAO struct {
	db *sql.DB
}

func NewFollowDAO(db *sql.DB) *FollowDAO {
	return &FollowDAO{db: db}
}

func (dao *FollowDAO) getTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value("currentTx").(*sql.Tx); ok {
		return tx
	}
	return nil
}

func (dao *FollowDAO) execContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return dao.db.ExecContext(ctx, query, args...)
}

func (dao *FollowDAO) queryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return dao.db.QueryRowContext(ctx, query, args...)
}

func (dao *FollowDAO) queryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if tx := dao.getTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return dao.db.QueryContext(ctx, query, args...)
}

func (dao *FollowDAO) Create(ctx context.Context, m *Follow) error {
	query := `
		INSERT INTO follows (id, follower_id, followee_id, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := dao.execContext(
		ctx,
		query,
		m.ID,
		m.FollowerID,
		m.FolloweeID,
		m.CreatedAt,
	)

	return err
}

func (dao *FollowDAO) Update(ctx context.Context, m *Follow) error {
	query := `
		UPDATE follows
		SET follower_id = $1,
			followee_id = $2,
			created_at = $3
		WHERE id = $4
	`

	_, err := dao.execContext(ctx, query,
		m.FollowerID,
		m.FolloweeID,
		m.CreatedAt,
		m.ID,
	)
	return err
}

func (dao *FollowDAO) PartialUpdate(ctx context.Context, pk string, fields map[string]interface{}) error {
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

	query := fmt.Sprintf(`UPDATE follows SET %s WHERE id = $%d`, strings.Join(setClauses, ", "), i)

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *FollowDAO) DeleteByPk(ctx context.Context, pk string) error {
	query := `DELETE FROM follows WHERE id = $1`
	_, err := dao.execContext(ctx, query, pk)
	return err
}

func (dao *FollowDAO) FindByPk(ctx context.Context, pk string) (*Follow, error) {
	query := `
		SELECT id, follower_id, followee_id, created_at
		FROM follows
		WHERE id = $1
	`
	row := dao.queryRowContext(ctx, query, pk)

	var m Follow
	err := row.Scan(
		&m.ID,
		&m.FollowerID,
		&m.FolloweeID,
		&m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *FollowDAO) CreateMany(ctx context.Context, models []*Follow) error {
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
			model.FollowerID,
			model.FolloweeID,
			model.CreatedAt,
		)
	}

	query := fmt.Sprintf(`
		INSERT INTO follows (id, follower_id, followee_id, created_at)
		VALUES %s
	`, strings.Join(placeholders, ", "))

	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *FollowDAO) UpdateMany(ctx context.Context, models []*Follow) error {
	if len(models) == 0 {
		return nil
	}

	query := `
		UPDATE follows
		SET follower_id = $1,
			followee_id = $2,
			created_at = $3
		WHERE id = $4
	`

	for _, model := range models {
		_, err := dao.execContext(ctx, query,
			model.FollowerID,
			model.FolloweeID,
			model.CreatedAt,
			model.ID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao *FollowDAO) DeleteManyByPks(ctx context.Context, pks []string) error {
	if len(pks) == 0 {
		return nil
	}

	placeholders := make([]string, len(pks))
	args := make([]interface{}, len(pks))
	for i, pk := range pks {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = pk
	}

	query := fmt.Sprintf(`DELETE FROM follows WHERE id IN (%s)`, strings.Join(placeholders, ","))
	_, err := dao.execContext(ctx, query, args...)
	return err
}

func (dao *FollowDAO) FindOne(ctx context.Context, where string, sort string, args ...interface{}) (*Follow, error) {
	query := `
		SELECT id, follower_id, followee_id, created_at
		FROM follows
	`

	if where != "" {
		query += " WHERE " + where
	}

	if sort != "" {
		query += " ORDER BY " + sort
	}

	row := dao.queryRowContext(ctx, query, args...)

	var m Follow
	err := row.Scan(
		&m.ID,
		&m.FollowerID,
		&m.FolloweeID,
		&m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (dao *FollowDAO) FindAll(ctx context.Context, where string, sort string, args ...interface{}) ([]*Follow, error) {
	query := `
		SELECT id, follower_id, followee_id, created_at
		FROM follows
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

	var models []*Follow
	for rows.Next() {
		var m Follow
		err := rows.Scan(
			&m.ID,
			&m.FollowerID,
			&m.FolloweeID,
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

func (dao *FollowDAO) FindPaginated(ctx context.Context, limit, offset int, where string, sort string, args ...interface{}) ([]*Follow, error) {
	query := `
		SELECT id, follower_id, followee_id, created_at
		FROM follows
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

	var models []*Follow
	for rows.Next() {
		var m Follow
		err := rows.Scan(
			&m.ID,
			&m.FollowerID,
			&m.FolloweeID,
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

func (dao *FollowDAO) Count(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := "SELECT COUNT(*) FROM follows"

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

func (dao *FollowDAO) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
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
