package copicake

import (
	"errors"
	"os"
	"time"
)

type Client struct {
	config *ClientConfig
}

type ClientConfig struct {
	ApiKey        string
	RetryTimeout  time.Duration
	RetryMaxTries int
}

// New returns a new Copicake client
func New(config *ClientConfig) (*Client, error) {

	// if config is nil, set default values
	if config == nil {
		config = &ClientConfig{
			RetryTimeout:  1 * time.Second,
			RetryMaxTries: 5,
			ApiKey:        os.Getenv("COPICAKE_API_KEY"),
		}
	}

	// vaildate client configs
	// if no API key is set, try to use the environment variable
	if config.ApiKey == "" {
		config.ApiKey = os.Getenv("COPICAKE_API_KEY")
	}

	// if still no API key, return error
	if config.ApiKey == "" {
		return nil, errors.New("copicake API key is required, either set it in the config or set the COPICAKE_API_KEY environment variable")
	}

	// if no retry timeout is set, set default value
	if config.RetryTimeout == 0 {
		config.RetryTimeout = 1 * time.Second
	}

	// if no retry max tries is set, set default value
	if config.RetryMaxTries == 0 {
		config.RetryMaxTries = 5
	}

	// return client
	return &Client{
		config: config,
	}, nil
}

// NewRenderRequest creates a new render job
func (c Client) NewRenderRequest(request RenderRequest) (*RenderJob, error) {
	return &RenderJob{}, nil
}
