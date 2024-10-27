package main

import (
	"context"

	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func seedExerciseMuscleGroups(db *database.Queries) error {
	ctx := context.TODO()

	items := []struct {
		exerciseName string
		muscleGroups []string
	}{
		{"Pushups", []string{"Chest", "Core", "Shoulders"}},
		{"Squats", []string{"Legs", "Glutes", "Core"}},
		{"Planks", []string{"Core", "Shoulders"}},
		{"Lunges", []string{"Legs", "Glutes", "Core"}},
		{"Pullups", []string{"Back", "Arms", "Shoulders"}},
	}

	for _, item := range items {
		exercise, err := db.GetExerciseByName(ctx, item.exerciseName)
		if err != nil {
			return err
		}

		for _, muscleGroupName := range item.muscleGroups {
			category, err := db.GetMuscleGroupByName(ctx, muscleGroupName)
			if err != nil {
				return err
			}

			_, err = db.AddExerciseMuscleGroup(ctx,
				database.AddExerciseMuscleGroupParams{
					ExerciseID:    exercise.ID,
					MuscleGroupID: category.ID,
				})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func clearExerciseMuscleGroups(db *database.Queries) error {
	ctx := context.TODO()

	exercises, err := db.GetExercises(ctx)
	if err != nil {
		return err
	}

	for _, e := range exercises {
		muscleGroupIds, err := db.GetExerciseMuscleGroups(ctx, e.ID)
		if err != nil {
			return err
		}

		for _, mgId := range muscleGroupIds {
			err := db.RemoveExerciseMuscleGroup(ctx, database.RemoveExerciseMuscleGroupParams{
				ExerciseID:    e.ID,
				MuscleGroupID: mgId,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
