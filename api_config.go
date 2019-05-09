package main

// APIConfig which stores the api configuration
type APIConfig struct {
	ConcurrentConnections int
	url                   string
	method                string
	header                map[string]string
	duration              int
	body                  string
	timeOut               int
}

func newAPIConfig(
	goroutines int,
	url string,
	method string,
	header map[string]string,
	duration int,
	body string,
	timeOut int,
) *APIConfig {
	return &APIConfig{
		ConcurrentConnections: goroutines,
		url:      url,
		method:   method,
		header:   header,
		duration: duration,
		body:     body,
		timeOut:  timeOut,
	}
}
