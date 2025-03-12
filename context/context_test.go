package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	data := "dummy data"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		responseBody := response.Body.String()

		if responseBody != data {
			t.Errorf("got %q, want %q", responseBody, data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel the work if request is canceled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancelingCtx)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}

// SpyStore allows you to simulate a store and see how its used.
type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

// Fetch returns response after a short delay.
func (s *SpyStore) Fetch() string {
	time.Sleep(50 * time.Millisecond)
	return s.response
}

// Cancel will record the call.
func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}
