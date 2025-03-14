package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	score PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.score.GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	switch name {
	case "John":
		return "42"
	case "Alice":
		return "24"
	default:
		return ""
	}
}
