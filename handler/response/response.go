package response

import (
	"encoding/json"
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/logger"
)

var jsonContentType = "application/json"

func WriteJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", jsonContentType)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "couldn't convert to json", http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, status int, err error, msg string) {
	if err != nil {
		logger.Log.Error(err.Error())
	}
	http.Error(w, msg, status)
}
