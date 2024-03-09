package main

import (
	"log"
	"net/http"
)

// ApiClient struct encapsulates logic for interacting with configured API
type ApiClient struct {
	*Service
	baseURL        string
	healthcheckURL string
	errorLog       *log.Logger
	infoLog        *log.Logger

	*http.Client
}

type Option interface {
	apply(*ApiClient)
}

// Create types for your different options
type (
	errorLogOption loggerOption
	infoLogOption  loggerOption

	loggerOption struct {
		Log *log.Logger
	}
)

// apply errorLogOption
func (el errorLogOption) apply(client *ApiClient) {
	client.errorLog = el.Log
}

// apply errorLogOption
func (il infoLogOption) apply(client *ApiClient) {
	client.infoLog = il.Log
}

func New(service *Service, opts ...Option) *ApiClient {
	client := &ApiClient{
		Service:        service,
		baseURL:        service.baseURL,
		healthcheckURL: service.healthcheckURL,

		Client: &http.Client{},
	}

	return client.WithOptions(opts...)
}

func (c *ApiClient) WithOptions(opts ...Option) *ApiClient {
	client := c.clone()

	for _, o := range opts {
		o.apply(client)
	}

	return client
}

// WithErrorLogOption implements Option
func WithErrorLogOption(errorLog *log.Logger) Option {
	return errorLogOption{errorLog}
}

// WithInfoLogOption implements Option
func WithInfoLogOption(infoLog *log.Logger) Option {
	return infoLogOption{infoLog}
}

func (c *ApiClient) clone() *ApiClient {
	clone := &c
	return *clone
}

func (c *ApiClient) doHealthcheck() (bool, error) {
	result, err := c.Get(c.healthcheckURL)
	if err != nil {
		return false, err
	}

	return result.StatusCode == http.StatusOK, nil
}

// CallHealthcheck calls doHealthcheck and acts off the result to log whether the endpoint
// returned a 200 OK status.
func CallHealthcheck(client *ApiClient) {
	if healthy, err := client.doHealthcheck(); err != nil {
		client.errorLog.Printf("\nError: %v", err.Error())
	} else if !healthy {
		// Handle the case where the health check failed (if needed)
		// For example, you might want to perform some action when the health check is not healthy.
		client.infoLog.Printf("%s is not healthy.", client.healthcheckURL)
	}

	client.infoLog.Printf("%s is healthy.", client.healthcheckURL)
}
