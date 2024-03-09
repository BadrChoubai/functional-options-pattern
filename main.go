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
		errorLog: log.New(os.Stderr, "ERROR:\t", log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO:\t", log.Lshortfile),
	}

	// This example code uses the Uber API:
	// https://developer.uber.com/docs/riders/ride-requests/tutorials/api/introduction
	uberService := &service.Service{
		BaseURL:        "https://api.uber.com/",
		ClientId:       "https://api.uber.com/health",
		ClientSecret:   os.Getenv("UBER_API_CLIENT_SECRET"),
		HealthcheckURL: os.Getenv("UBER_API_HEALTHCHECK_URL"),
	}

	app.infoLog.Printf("creating api client for: %s", uberService.BaseURL)
	UberClient := client.New(
		client.WithServiceConnection(uberService),
		client.WithErrorLog(app.errorLog),
		client.WithInfoLog(app.infoLog),
	)

	client.CheckHealth(UberClient)
}
