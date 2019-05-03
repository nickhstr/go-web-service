package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Hello greets a given name.
func Hello(r *chi.Mux) {
	r.Get("/hello", hello)
	r.Get("/hello/{name:([a-z])+}", helloName)
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello!"))
	if err != nil {
		panic(err)
	}
}

func helloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	_, err := w.Write([]byte("Hello " + name + "!"))
	if err != nil {
		panic(err)
	}
}
