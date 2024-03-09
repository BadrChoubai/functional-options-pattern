package client

import (
	"github.com/badrchoubai/functional-options-example/internal/pkg/service"
	"log"
	"net/http"
	"os"
)

// ApiClient struct encapsulates logic for interacting with configured API provided by
// service.Service
type ApiClient struct {
	Service  *service.Service
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	*http.Client
}

// CheckHealth calls doHealthcheck and acts off the result to log whether the endpoint
// returned a 200 OK status.
func CheckHealth(client *ApiClient) {
	healthy, err := client.doHealthcheck()
	if err != nil {
		client.ErrorLog.Printf("Error: %v", err.Error())
		return
	}

	if healthy {
		client.InfoLog.Printf("%s is healthy.", client.Service.BaseURL)
	} else {
		client.InfoLog.Printf("%s is not healthy.", client.Service.BaseURL)
	}
}

func New(opts ...Option) *ApiClient {
	if opts == nil {
		return NewNop()
	}

	client := &ApiClient{
		Client:   &http.Client{},
		ErrorLog: log.New(os.Stderr, "ERROR: \t", log.LstdFlags),
		InfoLog:  log.New(os.Stdout, "INFO: \t", log.LstdFlags),
	}

	return client.WithOptions(opts...)
}

func NewNop() *ApiClient {
	return &ApiClient{
		Client:   &http.Client{},
		ErrorLog: log.New(os.Stderr, "ERROR: \t", log.LstdFlags),
		InfoLog:  log.New(os.Stdout, "INFO: \t", log.LstdFlags),
	}
}

func (c *ApiClient) WithOptions(opts ...Option) *ApiClient {
	client := c.clone()

	for _, o := range opts {
		o.apply(client)
	}

	return client
}

func (c *ApiClient) clone() *ApiClient {
	clone := &c
	return *clone
}

func (c *ApiClient) doHealthcheck() (bool, error) {
	result, err := c.Get(c.Service.HealthcheckURL)
	if err != nil {
		return false, err
	}

	return result.StatusCode == http.StatusOK, nil
}

// Get implemented to introduce logging inside of function call
func (c *ApiClient) Get(url string) (resp *http.Response, err error) {
	c.InfoLog.Printf("GET::%s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}
