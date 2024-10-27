package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func respondError(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(message+": %v", err)))
}

func parseQueryList(paramList string) ([]int32, error) {
	valueList := []int32{}
	if paramList == "" {
		return valueList, nil
	}

	for _, s := range strings.Split(paramList, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return []int32{}, err
		}

		valueList = append(valueList, int32(i))
	}

	return valueList, nil
}
