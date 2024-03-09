package client

import (
	"github.com/badrchoubai/functional-options-example/internal/pkg/service"
	"log"
	"net/http"
)

// ApiClient struct encapsulates logic for interacting with configured API
type ApiClient struct {
	errorLog *log.Logger
	infoLog  *log.Logger

	*http.Client
	service *service.Service
}

func New(service *service.Service, opts ...Option) *ApiClient {
	client := &ApiClient{
		service: service,
		Client:  &http.Client{},
	}

	return client.WithOptions(opts...)
}

func (c *ApiClient) clone() *ApiClient {
	clone := &c
	return *clone
}

func (c *ApiClient) WithOptions(opts ...Option) *ApiClient {
	client := c.clone()

	for _, o := range opts {
		o.apply(client)
	}

	return client
}

func (c *ApiClient) doHealthcheck() (bool, error) {
	result, err := c.Get(c.service.HealthcheckURL)
	if err != nil {
		return false, err
	}

	return result.StatusCode == http.StatusOK, nil
}

// CheckHealth calls doHealthcheck and acts off the result to log whether the endpoint
// returned a 200 OK status.
func CheckHealth(client *ApiClient) {
	if client.service.HealthcheckURL == "" {
		client.errorLog.Print("client has no healthcheck endpoint configured")
		return
	}

	healthy, err := client.doHealthcheck()
	if err != nil {
		client.errorLog.Printf("\nError: %v", err.Error())
		return
	}

	if healthy {
		client.infoLog.Printf("%s is healthy.", client.service.HealthcheckURL)
	} else {
		client.infoLog.Printf("%s is not healthy.", client.service.HealthcheckURL)
	}

}
