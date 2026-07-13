package startupaid

import (
	"context"
	"net/http"
)

// Audience is a mailing list.
type Audience struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListAudiences returns the account's mailing lists.
func (c *Client) ListAudiences(ctx context.Context) ([]Audience, error) {
	var out struct {
		Audiences []Audience `json:"audiences"`
	}
	if err := c.do(ctx, http.MethodGet, "/v1/audiences", nil, nil, &out); err != nil {
		return nil, err
	}
	return out.Audiences, nil
}

// AddContactRequest adds a contact to an audience.
type AddContactRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

// AddContact adds a contact to a mailing list.
func (c *Client) AddContact(ctx context.Context, audienceID string, req AddContactRequest) error {
	return c.do(ctx, http.MethodPost, "/v1/audiences/"+esc(audienceID)+"/contacts", nil, req, nil)
}
