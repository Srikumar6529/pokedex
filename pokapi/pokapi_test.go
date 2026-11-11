package pokapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestGetLocationAreasCachesResponses(t *testing.T) {
	var hits int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		fmt.Fprint(w, `{"count":2,"next":"https://example.com/next","previous":"https://example.com/prev","results":[{"name":"kanto-route-1","url":"https://example.com/1"}]}`)
	}))
	defer server.Close()

	client := NewClient(time.Minute)

	first, err := client.GetLocationAreas(server.URL)
	if err != nil {
		t.Fatalf("first request: %v", err)
	}
	second, err := client.GetLocationAreas(server.URL)
	if err != nil {
		t.Fatalf("second request: %v", err)
	}

	if got := atomic.LoadInt32(&hits); got != 1 {
		t.Fatalf("server hit %d times; want 1", got)
	}
	if first.Results[0].Name != second.Results[0].Name {
		t.Fatalf("cached response differs: %q vs %q", first.Results[0].Name, second.Results[0].Name)
	}
	if second.Next != "https://example.com/next" || second.Prev != "https://example.com/prev" {
		t.Fatalf("unexpected next/prev: %q, %q", second.Next, second.Prev)
	}
}

func TestGetLocationAreasBadStatus(t *testing.T) {
	var hits int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(time.Minute)

	if _, err := client.GetLocationAreas(server.URL); err == nil {
		t.Fatal("expected an error for a 500 response")
	}
	if _, err := client.GetLocationAreas(server.URL); err == nil {
		t.Fatal("expected an error for a 500 response")
	}
	if got := atomic.LoadInt32(&hits); got != 2 {
		t.Fatalf("server hit %d times; want 2 (errors must not be cached)", got)
	}
}

func TestGetLocationAreasInvalidJSON(t *testing.T) {
	var hits int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		fmt.Fprint(w, `not json`)
	}))
	defer server.Close()

	client := NewClient(time.Minute)

	if _, err := client.GetLocationAreas(server.URL); err == nil {
		t.Fatal("expected an error for invalid JSON")
	}
	if _, err := client.GetLocationAreas(server.URL); err == nil {
		t.Fatal("expected an error for invalid JSON")
	}
	if got := atomic.LoadInt32(&hits); got != 2 {
		t.Fatalf("server hit %d times; want 2 (errors must not be cached)", got)
	}
}
