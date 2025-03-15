package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
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

	// Create a unique hash
	h := md5.New()
	h.Write([]byte(requestPayload.Link))
	bs := h.Sum(nil)

	// Encode the hash
	uEnc := base64.RawURLEncoding.EncodeToString([]byte(bs))

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "short URL created successfully",
		Data:    fmt.Sprintf("http://www.localhost%v/%v", app.config.addr, uEnc),
	}, nil)
}

func (app *application) redirectToURL(w http.ResponseWriter, r *http.Request) {

	app.writeJSON(w, http.StatusBadGateway, jsonResponse{
		Error:   false,
		Message: "Successfully redirected.",
	}, nil)
}

func isValidURL(url string) bool {
	regex := `^(https?:\/\/)?(www\.)?[\w-]+(\.[\w-]+)+([\/?#][^\s]*)?$`
	re := regexp.MustCompile(regex)
	return re.MatchString(url)
}
