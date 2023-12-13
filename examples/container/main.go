package main

import (
	"log"
	"net/http"

	"github.com/bilus/microwire/container"
	"github.com/bilus/microwire/examples/container/templates"
)

var c, _ = container.New(
	container.Service{Name: "Hello", Path: "/apps/hello", Host: "localhost:8000"},
	container.Service{Name: "Goodbye", Path: "/apps/goodbye", Host: "localhost:8001"},
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if c.ProxyTurbo(w, r) {
		return
	}
	t := templates.ContainerTemplate(r.URL.String(), c.Services())
	err := t.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HandleRequest)
	err := http.ListenAndServe(":80", nil) //nolint: gosec
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
