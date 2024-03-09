package client

import "log"

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

// WithErrorLogOption implements Option
func WithErrorLogOption(errorLog *log.Logger) Option {
	return errorLogOption{errorLog}
}

// WithInfoLogOption implements Option
func WithInfoLogOption(infoLog *log.Logger) Option {
	return infoLogOption{infoLog}
}
