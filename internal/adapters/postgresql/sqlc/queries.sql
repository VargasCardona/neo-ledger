-- name: ListGuitars :many
SELECT * FROM guitars;

-- name: FindGuitarByID :one
SELECT * FROM guitars WHERE id = $1;

-- name: CreateGuitar :one
INSERT INTO guitars (
    brand,
    model,
    year,
    notes
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateGuitar :one
UPDATE guitars
SET
    brand      = $2,
    model      = $3,
    year       = $4,
    notes      = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteGuitar :exec
DELETE FROM guitars WHERE id = $1;
