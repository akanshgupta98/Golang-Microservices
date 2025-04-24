package main

import (
	"log"
	"logger-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// recieved JSON into var.
	log.Println("Write log handler called!!")
	var requestPayload JSONPayload

	_ = app.readJson(w, r, &requestPayload)

	// insert data

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	log.Println("Insert entry  called!!")

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	log.Println("Insert entry  completed!!")

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}
