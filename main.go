package main

import (
	"log"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	app := &application{
		errorLog: log.New(os.Stderr, "ERROR:\t", log.LstdFlags),
		infoLog:  log.New(os.Stdout, "INFO:\t", log.LstdFlags),
	}

	// This example code uses the Uber API:
	// https://developer.uber.com/docs/riders/ride-requests/tutorials/api/introduction
	uber := Service{
		baseURL:        os.Getenv("UBER_API_URL"),
		clientId:       os.Getenv("UBER_API_CLIENT_ID"),
		clientSecret:   os.Getenv("UBER_API_CLIENT_SECRET"),
		healthcheckURL: os.Getenv("UBER_API_HEALTHCHECK_URL"),
	}

	UberClient := New(
		&uber,
		WithErrorLogOption(app.errorLog),
		WithInfoLogOption(app.infoLog),
	)

	CallHealthcheck(UberClient)
}
