package main

// Service struct encapsulates configuration related to connecting to an API
type Service struct {
	baseURL        string
	clientId       string
	clientSecret   string
	healthcheckURL string
}
