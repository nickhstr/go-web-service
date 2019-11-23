package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Hello greets a given name.
func Hello(r *chi.Mux) {
	hr := helloRouter()
	r.Mount("/api", hr)
}

func helloRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/hello", hello)
	r.Get("/hello/{name:[a-z]+}", helloName)

	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello!"))
}

func helloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	_, _ = w.Write([]byte("Hello " + name + "!"))
}
