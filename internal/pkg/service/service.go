package service

// Service struct encapsulates configuration related to connecting to an API
type Service struct {
	BaseURL        string
	ClientID       string
	ClientSecret   string
	HealthcheckURL string
}
