package api

import "github.com/mosamadeeb/go-workout-tracker/internal/database"

type ServerState struct {
	DB *database.Queries
}
