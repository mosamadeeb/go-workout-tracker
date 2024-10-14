package api

import (
	"encoding/json"
	"net/http"
)

func HandleApi(mux *http.ServeMux, state ServerState) {
	handleApiExercises(mux, state)
}

func respondJson(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}