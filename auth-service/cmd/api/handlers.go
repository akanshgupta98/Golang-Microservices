package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user with db.
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		log.Println("Email is: ", requestPayload.Email)
		log.Println("error is: ", err.Error())
		log.Println("EMAIL FAILURE")
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return

	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		log.Println("PWD FAILURE")
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return

	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("LOgged in User: %s", requestPayload.Email),
		Data:    user,
	}
	app.writeJSON(w, http.StatusAccepted, payload)

}
