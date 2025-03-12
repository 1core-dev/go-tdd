package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store fetches data.
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server returns a handler for calling Store.
func Server(s Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.Fetch(r.Context())

		if err != nil {
			// TODO. Handle error
			return
		}
		fmt.Fprint(w, data)
	}
}
