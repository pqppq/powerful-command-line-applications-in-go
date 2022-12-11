/*
Copyright © 2022 pqppq

*/

package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testResp = map[string]struct {
	Status int
	Body   string
}{
	"resultsMany": {
		Status: http.StatusOK,
		Body: `{
  "results": [
    {
      "Task": "Task 1",
      "Done": false,
      "CreatedAt": "2019-10-28T08:23:38.310097076-04:00",
      "CompletedAt": "0001-01-01T00:00:00Z"
    },
    {
      "Task": "Task 2",
      "Done": false,
      "CreatedAt": "2019-10-28T08:23:38.323447798-04:00",
      "CompletedAt": "0001-01-01T00:00:00Z"
    }
  ],
  "date": 1572265440,
  "total_results": 2
}`,
	},
	"resultsOne": {
		Status: http.StatusOK,
		Body: `{
  "results": [
    {
      "Task": "Task 1",
      "Done": false,
      "CreatedAt": "2019-10-28T08:23:38.310097076-04:00",
      "CompletedAt": "0001-01-01T00:00:00Z"
    }
  ],
  "date": 1572265440,
  "total_results": 1
}`,
	},
	"noResults": {
		Status: http.StatusOK,
		Body: `{
  "results": [],
  "date": 1572265440,
  "total_results": 0
}`,
	},
	"root": {
		Status: http.StatusOK,
		Body:   "There's an API here",
	},
	"notFound": {
		Status: http.StatusNotFound,
		Body:   "404 - not found",
	},
	"created": {
		Status: http.StatusCreated,
		Body:   "",
	},
	"noContent": {
		Status: http.StatusNoContent,
		Body:   "",
	},
}

func mockServer(h http.HandlerFunc) (string, func()) {
	ts := httptest.NewServer(h)

	return ts.URL, func() {
		ts.Close()
	}
}

func TestCompleteAction(t *testing.T) {
	expURLPath := "/todo/1"
	expMethod := http.MethodPatch
	expQuery := "complete"
	expOut := "Item number 1 marked as completed\n"
	arg := "1"

	// Instantiate a test server for Complete test
	url, cleanup := mockServer(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != expURLPath {
				t.Errorf("Expected path %q, got %q", expURLPath, r.URL.Path)
			}

			if r.Method != expMethod {
				t.Errorf("Expected method %q, got %q", expMethod, r.Method)
			}

			if _, ok := r.URL.Query()[expQuery]; !ok {
				t.Errorf("Expected query %q not found in URL", expQuery)
			}

			w.WriteHeader(testResp["noContent"].Status)
			fmt.Fprintln(w, testResp["noContent"].Body)
		})
	defer cleanup()

	// Execute Complete test
	var out bytes.Buffer

	if err := completeAction(&out, url, arg); err != nil {
		t.Fatalf("Expected no error, got %q.", err)
	}

	if expOut != out.String() {
		t.Errorf("Expected output %q, got %q", expOut, out.String())
	}
}

func TestDelAction(t *testing.T) {
	expURLPath := "/todo/1"
	expMethod := http.MethodDelete
	expOut := "Item number 1 deleted\n"
	arg := "1"

	// Instantiate a test server for Del test
	url, cleanup := mockServer(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != expURLPath {
				t.Errorf("Expected path %q, got %q", expURLPath, r.URL.Path)
			}

			if r.Method != expMethod {
				t.Errorf("Expected method %q, got %q", expMethod, r.Method)
			}

			w.WriteHeader(testResp["noContent"].Status)
			fmt.Fprintln(w, testResp["noContent"].Body)
		})
	defer cleanup()
	// Execute Del test
	var out bytes.Buffer

	if err := delAction(&out, url, arg); err != nil {
		t.Fatalf("Expected no error, got %q.", err)
	}

	if expOut != out.String() {
		t.Errorf("Expected output %q, got %q", expOut, out.String())
	}
}
