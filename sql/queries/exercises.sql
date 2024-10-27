-- name: AddExerciseCategory :one
INSERT INTO exercise_categories (exercise_id, category_id)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: RemoveExerciseCategory :exec
DELETE FROM exercise_categories
WHERE exercise_id = $1 AND category_id = $2;

-- name: GetExerciseCategories :many
SELECT category_id FROM exercise_categories
WHERE exercise_id = $1;

-- name: GetExercisesByCategories :many
-- Returns exercises that have ALL of the given categories
SELECT e.* FROM exercises e
JOIN exercise_categories ON id = exercise_id
WHERE category_id = ANY(sqlc.arg('category_ids')::int[]);

-- name: AddExerciseMuscleGroup :one
INSERT INTO exercise_muscle_groups (exercise_id, muscle_group_id)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: RemoveExerciseMuscleGroup :exec
DELETE FROM exercise_muscle_groups
WHERE exercise_id = $1 AND muscle_group_id = $2;

-- name: GetExerciseMuscleGroups :many
SELECT muscle_group_id FROM exercise_muscle_groups
WHERE exercise_id = $1;

-- name: GetExercisesByMuscleGroups :many
-- Returns exercises that have ALL of the given muscle groups
SELECT e.* FROM exercises e
JOIN exercise_muscle_groups ON id = exercise_id
WHERE muscle_group_id = ANY(sqlc.arg('muscle_group_ids')::int[]);

-- name: GetExercisesByCategoriesAndMuscleGroups :many
-- Returns exercises that have ALL of the given categories and muscle groups
-- If either only categories or muscle groups are needed, use the respective query instead of this one
SELECT e.* FROM exercises e
JOIN exercise_categories c ON e.id = c.exercise_id
JOIN exercise_muscle_groups mg ON e.id = mg.exercise_id
WHERE c.category_id = ANY(sqlc.arg('category_ids')::int[])
AND mg.muscle_group_id = ANY(sqlc.arg('muscle_group_ids')::int[]);
