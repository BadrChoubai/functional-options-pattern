package client

import (
	"github.com/badrchoubai/functional-options-example/internal/pkg/service"
	"log"
)

type Option interface {
	apply(*ApiClient)
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
func (el errorLogOption) apply(client *ApiClient) {
	client.ErrorLog = el.Log
}

// apply errorLogOption
func (il infoLogOption) apply(client *ApiClient) {
	client.InfoLog = il.Log
}

// apply errorLogOption
func (sc serviceConnectionOption) apply(client *ApiClient) {
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
