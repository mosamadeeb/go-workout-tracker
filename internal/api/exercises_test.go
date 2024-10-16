package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ApiExercises(t *testing.T) {
	state := prepTestState()
	handler := prepTestHandler(state, handleApiExercises)

	tests := []struct {
		name     string
		args     apiTestArgs
		expected apiTestExpected
	}{
		{
			"fetch exercise by ID",
			apiTestArgs{nil, *httptest.NewRequest(http.MethodGet, "/api/exercises/1", nil)},
			apiTestExpected{
				http.StatusOK,
				"application/json",
				[]byte(`{
					"id": 1,
					"name": "Pushups",
					"description": "A pushup is a bodyweight exercise where you lower and raise your body using your arms while keeping your back straight."
				}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiTestHelper(t, handler, tt.args, tt.expected)
		})
	}
}
