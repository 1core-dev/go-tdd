package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	canceled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(50 * time.Microsecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "Hello, world"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		responseBody := response.Body.String()

		if responseBody != data {
			t.Errorf("got %q, want %q", responseBody, data)
		}

		if store.canceled {
			t.Error("it should not have cancelled the store")
		}
	})

	t.Run("tells store to cancel the work if request is canceled", func(t *testing.T) {
		data := "Dummy data"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancelingCtx)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !store.canceled {
			t.Error("store was not told to cancel")
		}
	})
}
