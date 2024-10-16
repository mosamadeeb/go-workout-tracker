package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func main() {
	clearFlag := flag.Bool("clear", false, "Clears the entire database.")

	flag.Parse()

	godotenv.Load()

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
	}

	queries := database.New(db)

	if *clearFlag {
		fmt.Println("The entire database will be cleared. Continue? (y/N)")

		var choice string
		fmt.Scan(&choice)

		if strings.ToLower(choice) != "y" {
			fmt.Println("Aborting.")
			return
		}

		fmt.Println("Clearing DB...")
		clearDB(queries)

		fmt.Println("Done!")
	} else {
		fmt.Println("Seeding DB...")
		seedDB(queries)

		fmt.Println("Done!")
	}
}

func seedDB(db *database.Queries) {
	// TODO: We should store errors and display them instead of aborting
	if err := seedExercises(db); err != nil {
		log.Fatal(err)
	}
}

func clearDB(db *database.Queries) {
	if err := clearExercises(db); err != nil {
		log.Fatal(err)
	}
}
