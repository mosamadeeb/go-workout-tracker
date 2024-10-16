package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
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
}

type ExerciseDto struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
