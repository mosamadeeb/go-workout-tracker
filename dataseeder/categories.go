package main

import (
	"context"

	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func seedCategories(db *database.Queries) error {
	categories := []string{
		"Strength",
		"Endurance",
		"Bodyweight",
		"Core",
		"Upper Body",
		"Lower Body",
	}

	for _, item := range categories {
		_, err := db.CreateCategory(context.TODO(), item)
		if err != nil {
			return err
		}
	}

	return nil
}

func clearCategories(db *database.Queries) error {
	categories, err := db.GetCategories(context.TODO())
	if err != nil {
		return err
	}

	for _, e := range categories {
		err = db.DeleteCategory(context.TODO(), e.ID)
		if err != nil {
			return err
		}
	}

	if err = db.ResetCategoryId(context.TODO()); err != nil {
		return err
	}

	return nil
}
