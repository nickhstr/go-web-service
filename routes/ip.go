package routes

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/nickhstr/goweb/dal/sling"
	"github.com/nickhstr/goweb/write"
)

var ipRegex = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)

// IP returns the server's IP address
func IP(r *mux.Router) {
	r.HandleFunc("/ip", ipHandler).Methods(http.MethodGet)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	s := sling.New().Get("http://checkip.dyndns.org")

	req, err := s.Request()
	if err != nil {
		write.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req = req.WithContext(r.Context())

	resp, err := s.Do(req, nil, nil)
	if err != nil {
		write.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		write.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ipAddress := ipRegex.Find(body)
	w.Write(ipAddress)
}
