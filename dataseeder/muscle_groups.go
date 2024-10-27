package main

import (
	"context"

	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func seedMuscleGroups(db *database.Queries) error {
	muscleGroups := []string{
		"Chest",
		"Back",
		"Arms",
		"Core",
		"Legs",
		"Glutes",
		"Shoulders",
	}

	for _, item := range muscleGroups {
		_, err := db.CreateMuscleGroup(context.TODO(), item)
		if err != nil {
			return err
		}
	}

	return nil
}

func clearMuscleGroups(db *database.Queries) error {
	muscleGroups, err := db.GetMuscleGroups(context.TODO())
	if err != nil {
		return err
	}

	for _, e := range muscleGroups {
		err = db.DeleteMuscleGroup(context.TODO(), e.ID)
		if err != nil {
			return err
		}
	}

	if err = db.ResetMuscleGroupId(context.TODO()); err != nil {
		return err
	}

	return nil
}
