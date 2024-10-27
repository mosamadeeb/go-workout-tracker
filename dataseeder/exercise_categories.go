package main

import (
	"context"

	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func seedExerciseCategories(db *database.Queries) error {
	ctx := context.TODO()

	items := []struct {
		exerciseName  string
		categoryNames []string
	}{
		{"Pushups", []string{"Strength", "Bodyweight", "Upper Body"}},
		{"Squats", []string{"Strength", "Lower Body", "Bodyweight"}},
		{"Planks", []string{"Core", "Endurance", "Bodyweight"}},
		{"Lunges", []string{"Strength", "Lower Body", "Bodyweight"}},
		{"Pullups", []string{"Strength", "Upper Body", "Bodyweight"}},
	}

	for _, item := range items {
		exercise, err := db.GetExerciseByName(ctx, item.exerciseName)
		if err != nil {
			return err
		}

		for _, categoryName := range item.categoryNames {
			category, err := db.GetCategoryByName(ctx, categoryName)
			if err != nil {
				return err
			}

			_, err = db.AddExerciseCategory(ctx,
				database.AddExerciseCategoryParams{
					ExerciseID: exercise.ID,
					CategoryID: category.ID,
				})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func clearExerciseCategories(db *database.Queries) error {
	ctx := context.TODO()

	exercises, err := db.GetExercises(ctx)
	if err != nil {
		return err
	}

	for _, e := range exercises {
		categoryIds, err := db.GetExerciseCategories(ctx, e.ID)
		if err != nil {
			return err
		}

		for _, cId := range categoryIds {
			err := db.RemoveExerciseCategory(ctx, database.RemoveExerciseCategoryParams{
				ExerciseID: e.ID,
				CategoryID: cId,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
