package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

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

type apiTestArgs struct {
	w *httptest.ResponseRecorder
	r http.Request
}

type apiTestExpected struct {
	code        int
	contentType string
	body        []byte
}

func apiTestHelper(t *testing.T, handler http.Handler, args apiTestArgs, expected apiTestExpected) {
	t.Helper()

	if args.w == nil {
		args.w = httptest.NewRecorder()
	}

	handler.ServeHTTP(args.w, &args.r)

	res := args.w.Result()
	defer res.Body.Close()

	if res.StatusCode != expected.code {
		t.Logf("Expected status code %d but got %d\n", expected.code, res.StatusCode)

		if message, ok := checkApiError(res); ok {
			t.Logf("Error message: %s\n", message)
		}

		t.FailNow()
	}

	contentType := res.Header.Get("Content-Type")
	if expected.contentType != "" && contentType != expected.contentType {
		t.Logf("Expected content type %s but got %s\n", expected.contentType, contentType)
		t.FailNow()
	}

	switch contentType {
	case "application/json":
		var wantStruct, gotStruct map[string]interface{}

		expectedJson := json.NewDecoder(bytes.NewReader(expected.body))
		resultJson := json.NewDecoder(res.Body)

		expectedJson.Decode(&wantStruct)
		resultJson.Decode(&gotStruct)

		if !reflect.DeepEqual(wantStruct, gotStruct) {
			t.Logf("Expected json body %v but got %v\n", wantStruct, gotStruct)
			t.FailNow()
		}
	}
}
