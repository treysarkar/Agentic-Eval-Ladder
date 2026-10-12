package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/treysarkar/Agentic-Eval-Ladder/internal/api"
	"github.com/treysarkar/Agentic-Eval-Ladder/internal/store"
)

// todo mirrors the JSOn the api will return. Tests keep their own copy
// so they depend only on the HTTP contract, never on internal types

type todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

type apiError struct{
	Error string `json:"error"`
}

// new testServer starts a fresh serv ice with an empty store.
func newTestServer(t *testing.T) *httptest.Server{
	t.helper()
	srv:= httptest.NewServer(api.NewRouter(store.New()))
	t.Cleanup(srv.Close)
	return srv
}

// do sends a requesdt and returns the resposne and its body
// body may be nil, a string (sent as-is), or nay value (sent as JSON)

func do(t *testing.T, srv *httptest.Server, method, path string, body any) (*http.Response, []byte){
	t.Helper()
	var reader io.reader
	switch v := body.(type) {
	case nil:
	case string:
		reader = strings.NewReader(v)
	default:
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("marshal body: %v", err)
		}
		reader = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, srv.URL+path, reader)
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("%s %s: %v", method, path, err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	return resp, data
}

// decode parses a JSON body into T, failing the test on bad JSON.
func decode[T any](t *testing.T, data []byte) T {
	t.Helper()
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatalf("decode %q: %v", data, err)
	}
	return v
}

func expectStatus(t *testing.T, resp *http.Response, body []byte, want int) {
	t.Helper()
	if resp.StatusCode != want {
		t.Fatalf("status = %d, want %d (body: %s)", resp.StatusCode, want, body)
	}
}
