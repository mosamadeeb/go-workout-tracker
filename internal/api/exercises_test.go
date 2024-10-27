package api

import (
	"encoding/json"
	"fmt"
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
				apiTestExpectedContent(
					"application/json",
					[]byte(`{
						"id": 1,
						"name": "Pushups",
						"description": "A pushup is a bodyweight exercise where you lower and raise your body using your arms while keeping your back straight."
					}`),
				),
			},
		},
		{
			"fetch exercises by category",
			// categories=4 is "Core"
			apiTestArgs{nil, *httptest.NewRequest(http.MethodGet, "/api/exercises?categories=4", nil)},
			apiTestExpected{
				http.StatusOK,
				func(r *http.Response) error {
					var exercises []ExerciseDto

					resultJson := json.NewDecoder(r.Body)
					resultJson.Decode(&exercises)

					found := false
					for _, e := range exercises {
						if e.Name == "Planks" {
							found = true
							break
						}
					}

					if !found {
						return fmt.Errorf("Expected exercises to include Planks")
					}

					return nil
				},
			},
		},
		{
			"fetch exercises by category and muscle groups",
			// categories=1,5 is "Strength" and "Upper Body"
			// muscle_groups=7 is "Shoulders"
			apiTestArgs{nil, *httptest.NewRequest(http.MethodGet, "/api/exercises?categories=1,5&muscle_groups=7", nil)},
			apiTestExpected{
				http.StatusOK,
				func(r *http.Response) error {
					var exercises []ExerciseDto

					resultJson := json.NewDecoder(r.Body)
					resultJson.Decode(&exercises)

					// Pushups and pullups have Strength and Upper Body categories and Shoulders muscle group
					toFind := map[string]bool{
						"Pushups": false,
						"Pullups": false,
					}

					// Planks has Strength and Shoulders, but does not have the Upper Body category
					toNotFind := map[string]bool{
						"Planks": false,
					}

					for _, e := range exercises {
						if _, ok := toFind[e.Name]; ok {
							toFind[e.Name] = true
						}

						if _, ok := toNotFind[e.Name]; ok {
							toNotFind[e.Name] = true
						}
					}

					notFoundSlice := []string{}
					for name, found := range toFind {
						if !found {
							notFoundSlice = append(notFoundSlice, name)
						}
					}

					if len(notFoundSlice) != 0 {
						return fmt.Errorf("Expected exercises to include %v", notFoundSlice)
					}

					foundSlice := []string{}
					for name, found := range toNotFind {
						if found {
							foundSlice = append(foundSlice, name)
						}
					}

					if len(foundSlice) != 0 {
						return fmt.Errorf("Expected exercises to NOT include %v", foundSlice)
					}

					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiTestHelper(t, handler, tt.args, tt.expected)
		})
	}
}
