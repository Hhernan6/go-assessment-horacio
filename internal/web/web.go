package web

import (
	"encoding/json"
	"net/http"
)

// Respond will setup the response with the data provided
func Respond(message interface{}, w http.ResponseWriter, statusCode int) {
	jsonResp, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RespondError(message string, w http.ResponseWriter, statusCode int) {
	Respond(Response{Error: message}, w, statusCode)
}
