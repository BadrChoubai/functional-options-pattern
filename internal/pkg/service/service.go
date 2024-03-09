package service

// Service struct encapsulates configuration related to connecting to an API
type Service struct {
	BaseURL        string
	ClientId       string
	ClientSecret   string
	HealthcheckURL string
}
