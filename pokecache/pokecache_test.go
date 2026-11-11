package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Minute)
	cache.Add("first", []byte("original"))
	cache.Add("second", []byte("other"))
	cache.Add("first", []byte("updated"))

	for _, tc := range []struct {
		key   string
		want  string
		found bool
	}{
		{key: "first", want: "updated", found: true},
		{key: "second", want: "other", found: true},
		{key: "missing", found: false},
	} {
		t.Run(tc.key, func(t *testing.T) {
			got, ok := cache.Get(tc.key)
			if ok != tc.found || string(got) != tc.want {
				t.Fatalf("Get(%q) = %q, %v; want %q, %v", tc.key, got, ok, tc.want, tc.found)
			}
			if !ok && got != nil {
				t.Fatal("missing key should return nil")
			}
		})
	}
}

func TestFreshEntryStays(t *testing.T) {
	cache := NewCache(50 * time.Millisecond)
	cache.Add("fresh", []byte("value"))

	time.Sleep(20 * time.Millisecond)

	got, ok := cache.Get("fresh")
	if !ok || string(got) != "value" {
		t.Fatalf("fresh entry missing: got %q, %v", got, ok)
	}
}

func TestOldEntryRemoved(t *testing.T) {
	interval := 20 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("old", []byte("value"))

	time.Sleep(2*interval + 20*time.Millisecond)

	got, ok := cache.Get("old")
	if ok {
		t.Fatalf("old entry should have been reaped, got %q", got)
	}
}
