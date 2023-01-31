package copicake

import "time"

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
	ID string `json:"id"`
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

// Status returns render job status
func (r *RenderJob) Status() (*RenderJobStatus, error) {
	return &RenderJobStatus{}, nil
}

// Image wait server finish render and return image content
func (r *RenderJob) Image() ([]byte, error) {
	return []byte{}, nil
}

// URL wait server finish render and return generated image url
func (r *RenderJob) URL() (string, error) {
	return "", nil
}
