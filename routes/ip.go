package routes

import (
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/nickhstr/goweb/dal"
)

var ipRegex = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)

// IP returns the server's IP address
func IP(r *chi.Mux) {
	r.Get("/ip", ipHandler)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	resp, err := dal.Get("http://checkip.dyndns.org")
	if err != nil {
		resp = []byte(err.Error())
		writeIPResponse(w, resp)
		return
	}

	ipAddress := ipRegex.Find(resp)
	writeIPResponse(w, ipAddress)
}

func writeIPResponse(w http.ResponseWriter, resp []byte) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(resp)
	if err != nil {
		panic(err)
	}
}
