-- name: ListGuitars :many
SELECT * FROM guitars;

-- name: FindGuitarByID :one
SELECT * FROM guitars WHERE id = $1;
