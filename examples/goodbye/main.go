package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bilus/microwire/examples/goodbye/templates"

	"github.com/bilus/microwire/service"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	m := fmt.Sprintf("Goodbye, the time is %s", time.Now().String())
	service.AddTurboStreamHeaders(w)
	t := templates.GoodbyeTemplate(m)
	err := t.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/apps/goodbye", HandleRequest)
	err := http.ListenAndServe(":8001", nil) //nolint: gosec
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
