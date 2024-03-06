package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Option interface {
	apply(*options)
}

// options struct with all of available options
type options struct {
	apiVersions   []string
	allowedScopes []string
}

// Create types for your different options
type apiVersionsOption []string
type allowedScopesOption []string

// apply apiVersionsOption
func (av apiVersionsOption) apply(opts *options) {
	opts.apiVersions = av
}

// apply allowedScopesOption
func (as allowedScopesOption) apply(opts *options) {
	opts.allowedScopes = as
}

// WithApiVersions implements Option
func WithApiVersions(versions []string) Option {
	return apiVersionsOption(versions)
}

// WithAllowedScopes implements Option
func WithAllowedScopes(scopes []string) Option {
	return allowedScopesOption(scopes)
}

func NewClient(url string, opts ...Option) (*ApiClient, error) {
	client := &ApiClient{
		apiUrl:     url,
		httpClient: &http.Client{},
	}

	for _, o := range opts {
		o.apply(&client.options)
	}

	return client, nil
}

type ApiClient struct {
	apiUrl     string
	httpClient *http.Client
	options    options
}

// CallHealthcheck performs a health check on the specified URL using a GET request.
// It returns true if the health check is successful (HTTP status code 200 OK),
// otherwise, it returns false and logs the error.
func (c ApiClient) CallHealthcheck(url string) (bool, error) {
	result, err := c.httpClient.Get(url)
	if err != nil {
		return false, err
	}

	return result.StatusCode == http.StatusOK, nil
}

func CheckHealth(clients ...*ApiClient) {
	for _, client := range clients {
		healthcheckUrl := fmt.Sprintf("https://%s/health", client.apiUrl)
		if healthy, err := client.CallHealthcheck(healthcheckUrl); err != nil {
			log.Printf("\nError calling healthcheck:\bURL: %s\nError Message: %s", healthcheckUrl, err.Error())
		} else if healthy {
			continue
		}
		// Handle the case where the health check failed (if needed)
		// For example, you might want to perform some action when the health check is not healthy.
		log.Printf("%s is not healthy.\n", healthcheckUrl)
	}
}

func (c *ApiClient) GetEnabledApiVersions() []string {
	return c.options.apiVersions
}

func (c *ApiClient) GetAllowedScopes() []string {
	return c.options.allowedScopes
}

func (c *ApiClient) String() string {
	return fmt.Sprintf("ApiClient{apiUrl: %s, allowedScopes: %v, apiVersions: %v}",
		c.apiUrl, c.options.allowedScopes, c.options.apiVersions)
}

func main() {
	serviceOneClient, err := NewClient(
		os.Getenv("SERVICE_ONE_BASE_URL"),
		WithApiVersions([]string{"2021-02-01"}),
		WithAllowedScopes([]string{"read:resource", "write:resource"}))

	if err != nil {
		fmt.Print("failed to create client: serviceOneClient")
	}

	fmt.Println(serviceOneClient.String())

	serviceTwoClient, err := NewClient(
		os.Getenv("SERVICE_TWO_BASE_URL"),
		WithApiVersions([]string{"2021-02-01"}),
		WithAllowedScopes([]string{"write:resource"}))

	if err != nil {
		fmt.Print("failed to create client: serviceTwoClient")
	}

	fmt.Println(serviceTwoClient.String())

	serviceThreeClient, err := NewClient(
		os.Getenv("SERVICE_THREE_BASE_URL"),
		WithApiVersions([]string{"2021-02-01"}),
		WithAllowedScopes([]string{"read:resource"}))

	if err != nil {
		fmt.Print("failed to create client: serviceThreeClient")
	}

	fmt.Println(serviceThreeClient.String())

	CheckHealth(
		serviceOneClient,
		serviceTwoClient,
		serviceThreeClient,
	)
}
