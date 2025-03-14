package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "OK",
	}, nil)
}
