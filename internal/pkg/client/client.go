package client

import (
	"log"
	"net/http"
	"os"

	"github.com/badrchoubai/functional-options-example/internal/pkg/service"
)

// APIClient struct encapsulates logic for interacting with configured API provided by
// service.Service
type APIClient struct {
	Service  *service.Service
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	*http.Client
}

// CheckHealth calls doHealthcheck and acts off the result to log whether the endpoint
// returned a 200 OK status.
func CheckHealth(client *APIClient) {
	healthy, err := client.doHealthcheck()
	if err != nil {
		client.ErrorLog.Printf("error: %v", err.Error())
		return
	}

	if !healthy {
		client.InfoLog.Printf("%s is not healthy", client.Service.BaseURL)
	} else {
		client.InfoLog.Printf("%s is healthy", client.Service.BaseURL)
	}
}

func New(opts ...Option) *APIClient {
	client := &APIClient{
		Client:   &http.Client{},
		ErrorLog: log.New(os.Stderr, "ERROR: \t", log.LstdFlags),
		InfoLog:  log.New(os.Stdout, "INFO: \t", log.LstdFlags),
	}

	if opts != nil {
		return client.WithOptions(opts...)
	}
	return client
}

func (c *APIClient) WithOptions(opts ...Option) *APIClient {
	client := c.clone()

	for _, o := range opts {
		o.apply(client)
	}

	return client
}

func (c *APIClient) clone() *APIClient {
	clone := *c
	return &clone
}

func (c *APIClient) doHealthcheck() (bool, error) {
	result, err := c.Get(c.Service.HealthcheckURL)

	// Close the response body in case of an error or after reading the response
	defer func() {
		if result != nil {
			result.Body.Close()
		}
	}()

	if err != nil {
		return false, err
	}

	return result.StatusCode == http.StatusOK, nil
}

// Get implemented to introduce logging inside of function call
func (c *APIClient) Get(url string) (resp *http.Response, err error) {
	c.InfoLog.Printf("GET::%s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err = c.Do(req)

	// Close the response body in case of an error or after reading the response
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	return resp, err
}
