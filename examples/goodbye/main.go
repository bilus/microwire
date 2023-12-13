package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bilus/microwire/examples/goodbye/templates"
	"github.com/bilus/microwire/turbo"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	m := fmt.Sprintf("Goodbye, the time is %s", time.Now().String())
	turbo.Stream(
		turbo.Update("app-container", templates.Container(m)),
		turbo.Update("title", templates.Title())).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/apps/goodbye", HandleRequest)
	err := http.ListenAndServe(":8001", nil) //nolint: gosec
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
