package main

import (
	"net/http"
	"regexp"
)

type shortUrl struct {
	Link string `json:"link"`
}

func (app *application) createShortURL(w http.ResponseWriter, r *http.Request) {

	var requestPayload shortUrl

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest, "invalid request payload")
		return
	}

	// Check if it is a valid link
	if !(isValidURL(requestPayload.Link)) {
		app.writeJSON(w, http.StatusOK, jsonResponse{
			Error:   true,
			Message: "invalid link",
		}, nil)
		return
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "short URL created successfully",
		Data:    requestPayload,
	}, nil)
}

func (app *application) redirectToOriginalURL(w http.ResponseWriter, r *http.Request) {

}

func isValidURL(url string) bool {
	regex := `^(https?:\/\/)?(www\.)?[\w-]+(\.[\w-]+)+([\/?#][^\s]*)?$`
	re := regexp.MustCompile(regex)
	return re.MatchString(url)
}
