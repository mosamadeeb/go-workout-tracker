package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_ApiExercises(t *testing.T) {
	state := prepTestState()
	handler := prepTestHandler(state, handleApiExercises)

	type args struct {
		w *httptest.ResponseRecorder
		r http.Request
	}
	tests := []struct {
		name         string
		args         args
		expectedCode int
		expectedType string
		expectedBody string
	}{
		{
			"Fetch exercise by ID",
			args{nil, *httptest.NewRequest(http.MethodGet, "/api/exercises/1", nil)},
			http.StatusOK,
			"application/json",
			`{
				"id": 1,
				"name": "Pushups",
				"description": "A pushup is a bodyweight exercise where you lower and raise your body using your arms while keeping your back straight."
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.w == nil {
				tt.args.w = httptest.NewRecorder()
			}

			handler.ServeHTTP(tt.args.w, &tt.args.r)

			res := tt.args.w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedCode {
				t.Logf("Expected status code %d but got %d\n", tt.expectedCode, res.StatusCode)

				if message, ok := checkApiError(res); ok {
					t.Logf("Error message: %s\n", message)
				}

				t.FailNow()
			}

			contentType := res.Header.Get("Content-Type")
			if tt.expectedType != "" && contentType != tt.expectedType {
				t.Logf("Expected content type %s but got %s\n", tt.expectedType, contentType)
				t.FailNow()
			}

			switch contentType {
			case "application/json":
				var wantStruct, gotStruct map[string]interface{}

				expectedJson := json.NewDecoder(strings.NewReader(tt.expectedBody))
				resultJson := json.NewDecoder(res.Body)

				expectedJson.Decode(&wantStruct)
				resultJson.Decode(&gotStruct)

				if !reflect.DeepEqual(wantStruct, gotStruct) {
					t.Logf("Expected json body %v but got %v\n", wantStruct, gotStruct)
					t.FailNow()
				}
			}
		})
	}
}
