package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(s Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Cancel()
		fmt.Fprint(w, s.Fetch())
	}
}
