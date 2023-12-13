package main

import (
	"log"
	"net/http"

	"github.com/bilus/microwire/examples/hello/templates"
	"github.com/bilus/microwire/turbo"

	"github.com/bilus/microwire/service"
)

type Form struct {
	FirstName string
}

var form Form // Don't do that!

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	stream := turbo.Stream(
		turbo.Update("app-container", templates.Form(form.FirstName)),
		turbo.Update("title", templates.Title(form.FirstName)),
	)
	stream.ServeHTTP(w, r)
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	service.AddTurboStreamHeaders(w)
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	form.FirstName = r.PostFormValue("fname")
	stream := turbo.Stream(
		turbo.Update("app-container", templates.Form(form.FirstName)),
		turbo.Update("title", templates.Title(form.FirstName)),
		turbo.Update("alert", templates.Alert(form.FirstName)),
	)
	stream.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/apps/hello/say", HandleForm)
	http.HandleFunc("/apps/hello", HandleRequest)
	err := http.ListenAndServe(":8002", nil) //nolint: gosec
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
