package routes

import (
	"net/http"
	"net/url"
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
	var (
		fc  dal.FetchConfig
		err error
	)

	fc.URL = url.URL{
		Scheme: "http",
		Host:   "checkip.dyndns.org",
	}

	resp, err := dal.Fetch(fc)
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
