package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "John" {
		fmt.Fprint(w, "42")
		return
	}

	if player == "Alice" {
		fmt.Fprint(w, "24")
		return
	}
}
