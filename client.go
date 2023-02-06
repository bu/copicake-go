package copicake

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

	// validate client configs
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
	// body
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// call API
	resp, err := c.call("POST", "image/create", &jsonBody)

	if err != nil {
		return nil, err
	}

	// parse response
	renderResp := new(RenderJobCreateResponse)
	err = json.Unmarshal(resp, renderResp)

	if err != nil {
		return nil, err
	}

	if renderResp.Error != "" {
		return nil, errors.New(renderResp.Error)
	}

	// create new Render job
	renderJob := &RenderJob{
		client: &c,
		ID:     renderResp.Data.ID,
	}

	return renderJob, nil
}

// post sends a POST request to the Copicake API
func (c Client) call(method string, path string, body *[]byte) ([]byte, error) {
	// create new client
	client := &http.Client{}

	// create new request
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(*body)
	}

	req, err := http.NewRequest(method, "https://api.copicake.com/v1/"+path, reader)
	req.Header.Add("Authorization", "Bearer "+c.config.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// close response
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// return response
	return respBody, nil
}
