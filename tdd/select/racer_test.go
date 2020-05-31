package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Naive", func(t *testing.T) {
		t.Skip()
		slowURL := "http://www.facebook.com"
		fastURL := "http://www.quii.co.uk"

		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("Mock server", func(t *testing.T) {
		slowServer := makeDelayedServer(200 * time.Millisecond)
		fastServer := makeDelayedServer(20 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("Concurrent by Select", func(t *testing.T) {
		slowServer := makeDelayedServer(200 * time.Millisecond)
		fastServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		timedOut := 10 * time.Second
		got, _ := ConcurrentRacer(slowURL, fastURL, timedOut)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 1s", func(t *testing.T) {
		timedOut := 10 * time.Millisecond

		server := makeDelayedServer(50 * time.Millisecond)
		defer server.Close()

		_, err := ConcurrentRacer(server.URL, server.URL, timedOut)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
