package main

import (
	"context"

	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func seedExercises(db *database.Queries) error {
	data := []database.CreateExerciseParams{
		{Name: "Pushups", Description: "A pushup is a bodyweight exercise where you lower and raise your body using your arms while keeping your back straight."},
		{Name: "Squats", Description: "A squat involves lowering your hips from a standing position and then returning to standing, engaging your legs and glutes."},
		{Name: "Planks", Description: "A plank is a core exercise where you hold a position similar to a pushup but with your weight resting on your forearms."},
		{Name: "Lunges", Description: "A lunge involves stepping forward with one leg, lowering your hips until both knees are bent at a 90-degree angle, then returning to the starting position."},
		{Name: "Pullups", Description: "A pullup is an upper-body exercise where you hang from a bar and pull yourself up until your chin is above the bar."},
	}

	for _, item := range data {
		_, err := db.CreateExercise(context.TODO(), item)
		if err != nil {
			return err
		}
	}

	return nil
}

func clearExercises(db *database.Queries) error {
	exercises, err := db.GetExercises(context.TODO())
	if err != nil {
		return err
	}

	for _, e := range exercises {
		err = db.DeleteExercise(context.TODO(), e.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
