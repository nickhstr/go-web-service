package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Hello greets a given name.
func Hello(r *mux.Router) {
	hr := r.PathPrefix("/api/hello").Subrouter()
	hr.HandleFunc("", hello).Methods(http.MethodGet)
	hr.HandleFunc("/{name:[a-z]+}", helloName).Methods(http.MethodGet)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func helloName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Fprintf(w, "Hello %s!\n", name)
}
