package copicake

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type RenderRequest struct {
	TemplateID string         `json:"template_id"`
	Changes    []C            `json:"changes,omitempty"`
	Options    *RenderOptions `json:"options,omitempty"`
}

type RenderOptions struct {
	WebhookURL string `json:"webhook_url,omitempty"`
}

type C map[string]string

type RenderJob struct {
	ID     string `json:"id"`
	client *Client
}

type RenderJobStatus struct {
	Type       string         `json:"type"`
	Status     string         `json:"status"`
	Changes    []C            `json:"changes,omitempty"`
	Options    *RenderOptions `json:"options,omitempty"`
	TemplateID string         `json:"template_id"`
	ImageURL   string         `json:"permanent_url,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	CreatedBy  string         `json:"created_by"`
	ID         string         `json:"id"`
}

type RenderJobCreateResponse struct {
	Error string          `json:"error,omitempty"`
	Data  RenderJobStatus `json:"data"`
}

// Status returns render job status
func (r *RenderJob) Status() (*RenderJobStatus, error) {
	// call api
	resp, err := r.client.call("GET", "image/get?id="+r.ID, nil)
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

	// return status
	return &renderResp.Data, nil
}

// Image wait server finish render and return image content
func (r *RenderJob) Image() ([]byte, error) {
	url, err := r.URL()
	if err != nil {
		return []byte{}, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != 200 {
		return []byte{}, errors.New("status code: " + strconv.Itoa(resp.StatusCode))
	}

	err = resp.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	return io.ReadAll(resp.Body)
}

// URL wait server finish render and return generated image url
func (r *RenderJob) URL() (string, error) {

	retryTimeout := r.client.config.RetryTimeout
	retryMaxTries := r.client.config.RetryMaxTries

	for i := 0; i < retryMaxTries; i++ {
		status, err := r.Status()
		if err != nil {
			return "", err
		}

		if status.Status == "success" {
			return status.ImageURL, nil
		}

		time.Sleep(retryTimeout)
	}

	return "", errors.New("timeout")
}
