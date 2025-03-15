package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type createURLRequest struct {
	Link string `json:"link"`
}

func (app *application) createShortURL(w http.ResponseWriter, r *http.Request) {

	var requestPayload createURLRequest

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
	checksumHex := fmt.Sprintf("%x", bs)

	ctx := context.Background()

	// Validate checksum
	exists, err := app.store.URLS.ValidateChecksum(ctx, checksumHex)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError, err.Error())
		return
	}

	// if a URL already exists then just return
	if exists {
		// TODO redirect to an existing URL
		return
	}

	// Get new row number for the new URL
	id, err := app.store.URLS.CreateShortURL(ctx, requestPayload.Link, checksumHex)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError, err.Error())
		return
	}

	// Encode the id base62
	shortCode := base64.RawStdEncoding.EncodeToString([]byte(strconv.FormatInt(id, 10)))

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "short URL created successfully",
		Data:    fmt.Sprintf("http://www.localhost%v/%v", app.config.addr, shortCode),
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
