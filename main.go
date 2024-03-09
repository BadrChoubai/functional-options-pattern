package main

import (
	"log"
	"os"

	"github.com/badrchoubai/functional-options-example/internal/pkg/client"
	"github.com/badrchoubai/functional-options-example/internal/pkg/service"
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
	uberService := &service.Service{
		BaseURL:        "https://api.uber.com/",
		ClientID:       os.Getenv("UBER_API_CLIENT_ID"),
		ClientSecret:   os.Getenv("UBER_API_CLIENT_SECRET"),
		HealthcheckURL: "https://api.uber.com/health",
	}

	app.infoLog.Printf("creating api client for: %s", uberService.BaseURL)
	UberClient := client.New(
		client.WithServiceConnection(uberService),
	)

	client.CheckHealth(UberClient)
}
