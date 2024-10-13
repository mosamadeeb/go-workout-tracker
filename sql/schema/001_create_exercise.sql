-- +goose Up
CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE muscle_groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Junction tables (many-to-many relations)
CREATE TABLE exercise_categories (
    exercise_id INT REFERENCES exercises,
    category_id INT REFERENCES categories,
    PRIMARY KEY (exercise_id, category_id)
);

CREATE TABLE exercise_muscle_groups (
    exercise_id INT REFERENCES exercises,
    muscle_group_id INT REFERENCES muscle_groups,
    PRIMARY KEY (exercise_id, muscle_group_id)
);

-- +goose Down
DROP TABLE exercise_muscle_groups;
DROP TABLE exercise_categories;
DROP TABLE muscle_groups;
DROP TABLE categories;
DROP TABLE exercises;
