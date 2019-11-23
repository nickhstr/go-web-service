package routes

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
)

func Debug(r *chi.Mux) {
	dr := debugRouter()
	r.Mount("/debug", dr)
}

func debugRouter() http.Handler {
	r := chi.NewRouter()
	// Redirect to /debug/pprof/, as some links in pprof's Index template
	// assume a trailing `/` in the URL
	r.HandleFunc("/pprof", func(w http.ResponseWriter, r *http.Request) {
		redirPath := r.URL.Path + "/"
		if r.URL.RawQuery != "" {
			redirPath = redirPath + "?" + r.URL.RawQuery
		}

		http.Redirect(w, r, redirPath, http.StatusMovedPermanently)
	})
	r.HandleFunc("/pprof/", pprof.Index)
	r.HandleFunc("/pprof/{subroute}", pprof.Index)
	r.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/pprof/profile", pprof.Profile)
	r.HandleFunc("/pprof/symbold", pprof.Symbol)
	r.HandleFunc("/pprof/trace", pprof.Trace)

	return r
}
