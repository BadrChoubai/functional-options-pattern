package client

import (
	"log"

	"github.com/badrchoubai/functional-options-example/internal/pkg/service"
)

type Option interface {
	apply(*APIClient)
}

// Create types for your different options
type (
	errorLogOption          logger
	infoLogOption           logger
	serviceConnectionOption serviceConnection

	serviceConnection struct {
		Service *service.Service
	}

	logger struct {
		Log *log.Logger
	}
)

// apply errorLogOption
func (el errorLogOption) apply(client *APIClient) {
	client.ErrorLog = el.Log
}

// apply errorLogOption
func (il infoLogOption) apply(client *APIClient) {
	client.InfoLog = il.Log
}

// apply errorLogOption
func (sc serviceConnectionOption) apply(client *APIClient) {
	client.Service = sc.Service
}

// WithErrorLog implements Option
func WithErrorLog(errorLog *log.Logger) Option {
	return errorLogOption{errorLog}
}

// WithInfoLog implements Option
func WithInfoLog(infoLog *log.Logger) Option {
	return infoLogOption{infoLog}
}

// WithServiceConnection implements Option
func WithServiceConnection(service *service.Service) Option {
	return serviceConnectionOption{
		Service: service,
	}
}
