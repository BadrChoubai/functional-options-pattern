package main

import (
	"fmt"
	"log"
	"net/http"
)

type Option interface {
	apply(*ApiClient)
}

// Create types for your different options
type (
	healthcheckURLOption string
	errorLogOption       loggerOption
	infoLogOption        loggerOption

	loggerOption struct {
		Log *log.Logger
	}
)

// apply healthcheckOption
func (hc healthcheckURLOption) apply(client *ApiClient) {
	client.healthcheckURL = string(hc)
}

// apply errorLogOption
func (el errorLogOption) apply(client *ApiClient) {
	client.errorLog = el.Log
}

// apply errorLogOption
func (il infoLogOption) apply(client *ApiClient) {
	client.infoLog = il.Log
}

type ApiClient struct {
	healthcheckURL string
	httpClient     *http.Client

	errorLog *log.Logger
	infoLog  *log.Logger
}

// WithHealthcheckURLOption implements Option
func WithHealthcheckURLOption(url string) Option {
	return healthcheckURLOption(url)
}

// WithErrorLogOption implements Option
func WithErrorLogOption(errorLog *log.Logger) Option {
	return errorLogOption{errorLog}
}

// WithInfoLogOption implements Option
func WithInfoLogOption(infoLog *log.Logger) Option {
	return infoLogOption{infoLog}
}

func New(opts ...Option) (*ApiClient, error) {
	client := &ApiClient{
		httpClient: &http.Client{},
	}

	for _, o := range opts {
		o.apply(client)
	}

	return client, nil
}

func (c ApiClient) doHealthcheck() (bool, error) {
	result, err := c.httpClient.Get(c.healthcheckURL)
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

func (c ApiClient) GetHealthcheckURL() string {
	return c.healthcheckURL
}

func (c ApiClient) String() string {
	return fmt.Sprintf("ApiClient{healthcheckURL: %s}",
		c.healthcheckURL)
}
