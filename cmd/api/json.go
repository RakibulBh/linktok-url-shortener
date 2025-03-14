package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status int, message string) error {

	var payload jsonResponse
	payload.Error = true
	payload.Message = message

	return app.writeJSON(w, status, payload, nil)

}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	return dec.Decode(data)
}
