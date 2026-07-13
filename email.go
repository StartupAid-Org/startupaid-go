package startupaid

import (
	"context"
	"net/http"
)

// SendEmailRequest is the payload for SendEmail. Provide either Body (HTML) or a
// Template name with Variables.
type SendEmailRequest struct {
	From      string            `json:"from"` // verified sender, e.g. "Acme <hi@acme.com>"
	To        string            `json:"to"`   // recipient address
	Subject   string            `json:"subject,omitempty"`
	Body      string            `json:"body,omitempty"`     // HTML or text body
	Template  string            `json:"template,omitempty"` // saved template name instead of Body
	Variables map[string]string `json:"variables,omitempty"`
}

// Message is an email record returned by SendEmail and GetMessage.
type Message struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Status    string `json:"status"`
	Template  string `json:"template,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

// SendEmail sends a transactional email.
func (c *Client) SendEmail(ctx context.Context, req SendEmailRequest) (*Message, error) {
	var m Message
	if err := c.do(ctx, http.MethodPost, "/v1/send", nil, req, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMessage returns a previously sent email's status by id.
func (c *Client) GetMessage(ctx context.Context, id string) (*Message, error) {
	var m Message
	if err := c.do(ctx, http.MethodGet, "/v1/messages/"+esc(id), nil, nil, &m); err != nil {
		return nil, err
	}
	return &m, nil
}
