package handlers

import (
	"fmt"
	"net/http"

	"github.com/nickhstr/go-web-service/app/utils"
)

// Health reports general information about the service
func Health(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s OK", utils.AppName())
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
