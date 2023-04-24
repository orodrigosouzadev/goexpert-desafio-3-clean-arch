-- name: ListOrders :many
SELECT * FROM orders;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = ?;

-- name: InsertOrder :exec
INSERT INTO orders (id, price, tax, final_price)
VALUES (?, ?, ?, ?);
