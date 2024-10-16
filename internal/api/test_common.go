package api

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func prepTestState() ServerState {
	godotenv.Load("../../.env")
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
	}

	return ServerState{DB: database.New(db), IsDev: true}
}

func prepTestHandler(state ServerState, handleApiFunc func(mux *http.ServeMux, state ServerState)) http.Handler {
	mux := http.NewServeMux()
	handleApiFunc(mux, state)
	return mux
}

func checkApiError(r *http.Response) (string, bool) {
	if r.Header.Get("Content-Type") == "text/plain" {
		defer r.Body.Close()
		message, _ := io.ReadAll(r.Body)
		return string(message), true
	}

	return "", false
}
