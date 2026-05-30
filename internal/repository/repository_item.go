package repository

import (
	"context"
	"errors"
	"fmt"

	"job4j.ru/go-lang-base/internal/tracker"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool: pool}
}

func (r *RepoPg) Pool() *pgxpool.Pool {
	return r.pool
}

func (r *RepoPg) Create(ctx context.Context, it tracker.Item) error {
	_, err := r.pool.Exec(
		ctx,
		`insert into items(id, name) values($1, $2)`,
		it.ID, it.Name,
	)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}
	return nil
}

func (r *RepoPg) List(ctx context.Context) ([]tracker.Item, error) {
	rows, err := r.pool.Query(ctx, `select id, name from items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var item tracker.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepoPg) Get(ctx context.Context, id string) (tracker.Item, error) {
	var query = `SELECT id, name FROM items WHERE id = $1`
	var item tracker.Item
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&item.ID,
		&item.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return tracker.Item{}, tracker.ErrNotFound
		}
		return tracker.Item{}, fmt.Errorf("r.pool.QueryRow: %w", err)
	}

	return item, nil
}

func (r *RepoPg) FindByName(ctx context.Context, name string) ([]tracker.Item, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, name FROM items WHERE name = $1`, name)
	if err != nil {
		return nil, fmt.Errorf("r.pool.Query: %w", err)
	}
	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var item tracker.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return items, nil
}

func (r *RepoPg) Update(ctx context.Context, item tracker.Item) (tracker.Item, error) {
	var query = `UPDATE items SET name = $2 WHERE id = $1 RETURNING id, name`
	var updated tracker.Item
	err := r.pool.QueryRow(ctx, query, item.ID, item.Name).Scan(
		&updated.ID,
		&updated.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return tracker.Item{}, tracker.ErrNotFound
		}
		return tracker.Item{}, fmt.Errorf("r.pool.QueryRow: %w", err)
	}
	return updated, nil
}

func (r *RepoPg) Delete(ctx context.Context, id string) error {
	var query = `DELETE FROM items WHERE id = $1`
	res, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}
	if res.RowsAffected() == 0 {
		return tracker.ErrNotFound
	}
	return nil
}
