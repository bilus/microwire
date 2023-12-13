package main

import (
	"log"
	"net/http"

	"github.com/bilus/microwire/examples/hello/templates"

	"github.com/bilus/microwire/service"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("HELLO", r.URL.String())
	service.AddTurboStreamHeaders(w)
	t := templates.FormTemplateStream("John")
	err := t.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	service.AddTurboStreamHeaders(w)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	t := templates.FormTemplateStream(r.PostFormValue("fname"))

	err = t.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/apps/hello/say", HandleForm)
	http.HandleFunc("/apps/hello", HandleRequest)
	err := http.ListenAndServe(":8000", nil) //nolint: gosec
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
