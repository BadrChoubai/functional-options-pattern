package main

import (
	"github.com/badrchoubai/functional-options-example/internal/pkg/client"
	"github.com/badrchoubai/functional-options-example/internal/pkg/service"
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
	uber := &service.Service{
		BaseURL:        os.Getenv("UBER_API_URL"),
		ClientId:       os.Getenv("UBER_API_CLIENT_ID"),
		ClientSecret:   os.Getenv("UBER_API_CLIENT_SECRET"),
		HealthcheckURL: os.Getenv("UBER_API_HEALTHCHECK_URL"),
	}

	UberClient := client.New(
		uber,
		client.WithErrorLogOption(app.errorLog),
		client.WithInfoLogOption(app.infoLog),
	)

	client.CheckHealth(UberClient)
}
