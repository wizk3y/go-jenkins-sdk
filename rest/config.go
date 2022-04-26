package restclient

import (
	"net/http"
	"os"
	"time"

	"github.com/wizk3y/go-jenkins-sdk/logger"
)

// Config --
type Config struct {
	BaseURL string

	Username string
	Password string

	Logger logger.LoggerInterface
}

// NewConfig --
func NewConfig(baseURL, username, password string) *Config {
	l := logger.GetDefaultLogger(os.Stdout)

	return &Config{
		BaseURL: baseURL,

		Username: username,
		Password: password,

		Logger: l,
	}
}

// Build --
func (c *Config) Build() (*RESTClient, error) {
	// prepare transport for new client
	r := http.DefaultTransport
	r = WrapWithBasicAuthRoundTripper(r, c.Username, c.Password)
	r = WrapWithLogRoundTripper(r, c.Logger)

	httpClient := http.Client{
		Timeout:   10 * time.Second,
		Transport: r,
	}

	return &RESTClient{
		baseURL: c.BaseURL,

		client: &httpClient,
	}, nil
}

// MustBuild --
func (c *Config) BuildOrDie() *RESTClient {
	// prepare transport for new client
	r := http.DefaultTransport
	r = WrapWithBasicAuthRoundTripper(r, c.Username, c.Password)
	r = WrapWithLogRoundTripper(r, c.Logger)

	httpClient := http.Client{
		Timeout:   10 * time.Second,
		Transport: r,
	}

	return &RESTClient{
		baseURL: c.BaseURL,

		client: &httpClient,
	}
}
