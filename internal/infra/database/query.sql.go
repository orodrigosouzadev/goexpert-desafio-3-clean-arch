// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package database

import (
	"context"
)

const getOrder = `-- name: GetOrder :one
SELECT id, price, tax, final_price FROM orders
WHERE id = ?
`

func (q *Queries) GetOrder(ctx context.Context, id string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Price,
		&i.Tax,
		&i.FinalPrice,
	)
	return i, err
}

const insertOrder = `-- name: InsertOrder :exec
INSERT INTO orders (id, price, tax, final_price)
VALUES (?, ?, ?, ?)
`

type InsertOrderParams struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) error {
	_, err := q.db.ExecContext(ctx, insertOrder,
		arg.ID,
		arg.Price,
		arg.Tax,
		arg.FinalPrice,
	)
	return err
}

const listOrders = `-- name: ListOrders :many
SELECT id, price, tax, final_price FROM orders
`

func (q *Queries) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Tax,
			&i.FinalPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
