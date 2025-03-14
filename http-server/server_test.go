package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayer(t *testing.T) {
	t.Run("returns John's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/John", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "42"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns Alice's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Alice", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "24"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
