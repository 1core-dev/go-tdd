package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Greet sends a personalized greeting to writer.
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

// GreeterHandler says Hello, world over HTTP.
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5050", http.HandlerFunc(GreetHandler)))
}
