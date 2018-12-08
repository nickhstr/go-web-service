package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Hello greets a given name.
func Hello() {
	Router.Get("/hello", hello)
	Router.Get("/hello/{name:([a-z])+}", helloName)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func helloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Write([]byte("Hello " + name + "!"))
}
