-- name: CreateExercise :one
INSERT INTO exercises (name, description)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetExercise :one
SELECT * FROM exercises
WHERE id = $1;

-- name: GetExercises :many
-- TODO: Add pagination/limits or remove this query
SELECT * FROM exercises;

-- name: UpdateExercise :exec
UPDATE exercises
SET name = $2,
    description = $3
WHERE id = $1;

-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1;

-- name: ResetExerciseId :exec
SELECT setval('exercises_id_seq', 1, FALSE);
