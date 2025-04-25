package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const webPort = "80"

func main() {
	// Load config
	app := Config{}

	// Start Server

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	log.Println("Mail service started on port: ", webPort)
	if err != nil {
		log.Panic(err)

	}

}
