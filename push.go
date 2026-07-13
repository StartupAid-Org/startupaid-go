package startupaid

import (
	"context"
	"net/http"
)

// PushTarget selects which devices a push goes to. Set exactly one of the fields.
type PushTarget struct {
	UserRef  string   `json:"userRef,omitempty"`
	UserRefs []string `json:"userRefs,omitempty"`
	Tokens   []string `json:"tokens,omitempty"`
	All      bool     `json:"all,omitempty"`
}

// SendPushRequest is the payload for SendPush.
type SendPushRequest struct {
	Target PushTarget        `json:"target"`
	Title  string            `json:"title,omitempty"`
	Body   string            `json:"body,omitempty"`
	Link   string            `json:"link,omitempty"`
	Data   map[string]string `json:"data,omitempty"`
}

// PushResult summarizes a push send.
type PushResult struct {
	MessageID  string `json:"messageId"`
	Status     string `json:"status"`
	Recipients int    `json:"recipients"`
	Sent       int    `json:"sent"`
	Failed     int    `json:"failed"`
}

// SendPush sends a push notification to devices in a push app.
func (c *Client) SendPush(ctx context.Context, appID string, req SendPushRequest) (*PushResult, error) {
	var r PushResult
	if err := c.do(ctx, http.MethodPost, "/v1/push/apps/"+esc(appID)+"/send", nil, req, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
