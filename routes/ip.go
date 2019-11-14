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
	w.Header().Set("Content-Type", "text/plain")

	resp, err := dal.Get("http://checkip.dyndns.org")
	if err != nil {
		resp = []byte(err.Error())
		_, _ = w.Write(resp)
		return
	}

	ipAddress := ipRegex.Find(resp)
	_, _ = w.Write(ipAddress)
}
