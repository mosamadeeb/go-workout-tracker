-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES (
    $1
)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1;

-- name: GetCategoryByName :one
SELECT * FROM categories
WHERE name = $1;

-- name: GetCategories :many
SELECT * FROM categories;

-- name: UpdateCategory :exec
UPDATE categories
SET name = $2
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: ResetCategoryId :exec
SELECT setval('categories_id_seq', 1, FALSE);
