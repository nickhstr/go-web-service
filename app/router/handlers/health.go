package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nickhstr/go-web-service/app/types"
	"github.com/nickhstr/go-web-service/app/utils"
)

// HealthInfo defines the set of info for the health check route
type HealthInfo struct {
	App    string  `json:"app"`
	Uptime float64 `json:"uptime"`
}

// Health reports general information about the service
func Health(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		res           []byte
		err           error
		transactionID string
	)

	transactionID, ok := r.Context().Value(types.TransactionKey).(string)
	if !ok {
		transactionID = "default"
	}

	appName := utils.App.Name()
	uptime := utils.App.Uptime()

	hi := HealthInfo{
		App:    appName,
		Uptime: uptime,
	}

	res, err = json.Marshal(hi)
	if err != nil {
		res, err = json.Marshal(types.NewError(transactionID, err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}
