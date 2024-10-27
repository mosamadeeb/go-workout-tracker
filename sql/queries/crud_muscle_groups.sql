-- name: CreateMuscleGroup :one
INSERT INTO muscle_groups (name)
VALUES (
    $1
)
RETURNING *;

-- name: GetMuscleGroup :one
SELECT * FROM muscle_groups
WHERE id = $1;

-- name: GetMuscleGroupByName :one
SELECT * FROM muscle_groups
WHERE name = $1;

-- name: GetMuscleGroups :many
SELECT * FROM muscle_groups;

-- name: UpdateMuscleGroup :exec
UPDATE muscle_groups
SET name = $2
WHERE id = $1;

-- name: DeleteMuscleGroup :exec
DELETE FROM muscle_groups
WHERE id = $1;

-- name: ResetMuscleGroupId :exec
SELECT setval('muscle_groups_id_seq', 1, FALSE);
