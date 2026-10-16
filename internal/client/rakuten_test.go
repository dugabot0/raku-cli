package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Rakuten explains rejected queries in the response body; dropping it leaves
// callers with a bare "API error (HTTP 400)" and no way to tell why.
func TestFetchIncludesErrorDescriptionOnAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"wrong_parameter","error_description":"NGKeyword is not valid"}`))
	}))
	defer srv.Close()

	c := New("app", "", "", "", 5*time.Second)
	_, err := c.fetch(srv.URL)
	if err == nil {
		t.Fatal("expected an error")
	}
	for _, want := range []string{"HTTP 400", "wrong_parameter", "NGKeyword is not valid"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not contain %q", err.Error(), want)
		}
	}
}

func TestFetchFallsBackToStatusWhenBodyIsNotJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<html>bad gateway</html>"))
	}))
	defer srv.Close()

	c := New("app", "", "", "", 5*time.Second)
	_, err := c.fetch(srv.URL)
	if err == nil || err.Error() != "API error (HTTP 502)" {
		t.Fatalf("got %v, want API error (HTTP 502)", err)
	}
}
