package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/mosamadeeb/go-workout-tracker/internal/api"
	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func main() {
	godotenv.Load()
	state := api.ServerState{}

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080", // TODO: load from .env and check dev/prod mode
		Handler: mux,
	}

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
	}

	state.DB = database.New(db)

	api.HandleApi(mux, state)

	fmt.Println("Starting server")
	server.ListenAndServe()
}
