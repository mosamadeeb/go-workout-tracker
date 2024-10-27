package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/mosamadeeb/go-workout-tracker/internal/database"
)

func handleApiExercises(mux *http.ServeMux, state ServerState) {
	mux.HandleFunc("GET /api/exercises/{exerciseId}", func(w http.ResponseWriter, r *http.Request) {
		exerciseId, err := strconv.Atoi(r.PathValue("exerciseId"))
		if err != nil {
			respondError(w, http.StatusBadRequest, "could not convert path value", err)
			return
		}

		exercise, err := state.DB.GetExercise(r.Context(), int32(exerciseId))
		if errors.Is(err, sql.ErrNoRows) {
			respondError(w, http.StatusNotFound, "exercise not found", err)
			return
		} else if err != nil {
			respondError(w, http.StatusInternalServerError, "database error", err)
			return
		}

		respondJson(w, http.StatusOK, ExerciseDto{
			exercise.ID,
			exercise.Name,
			exercise.Description,
		})
	})

	// Request format: GET /api/exercises?categories=1,3&muscle_groups=2
	mux.HandleFunc("GET /api/exercises", func(w http.ResponseWriter, r *http.Request) {
		// Check for categories and muscle groups in the path parameters
		query := r.URL.Query()
		categories, err := parseQueryList(query.Get("categories"))
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid query format (categories)", err)
			return
		}

		muscleGroups, err := parseQueryList(query.Get("muscle_groups"))
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid query format (muscle_groups)", err)
			return
		}

		var exercises []database.Exercise
		if len(categories) != 0 && len(muscleGroups) != 0 {
			// Both lists have values
			exercises, err = state.DB.GetExercisesByCategoriesAndMuscleGroups(r.Context(),
				database.GetExercisesByCategoriesAndMuscleGroupsParams{
					CategoryIds:    categories,
					MuscleGroupIds: muscleGroups,
				})
		} else {
			if len(categories) == 0 {
				// Only muscle groups in the query
				exercises, err = state.DB.GetExercisesByMuscleGroups(r.Context(), muscleGroups)
			} else if len(muscleGroups) == 0 {
				// Only categories in the query
				exercises, err = state.DB.GetExercisesByCategories(r.Context(), categories)
			} else {
				// No query parameters
				exercises, err = state.DB.GetExercises(r.Context())
			}
		}

		if errors.Is(err, sql.ErrNoRows) {
			respondError(w, http.StatusNotFound, "no exercises exist", err)
			return
		} else if err != nil {
			respondError(w, http.StatusInternalServerError, "database error", err)
			return
		}

		exercisesArr := make([]ExerciseDto, len(exercises))
		for i := range exercises {
			exercisesArr[i] = ExerciseDto{
				exercises[i].ID,
				exercises[i].Name,
				exercises[i].Description,
			}
		}

		respondJson(w, http.StatusOK, exercisesArr)
	})
}

type ExerciseDto struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
