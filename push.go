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

// RegisterDeviceRequest registers (or refreshes) a device token in a push app.
// Platform is "web", "ios", or "android". Pass a stable UserRef so you can later
// target the user with PushTarget{UserRef: ...}.
type RegisterDeviceRequest struct {
	Token      string `json:"token"`
	Platform   string `json:"platform"`
	UserRef    string `json:"userRef,omitempty"`
	Locale     string `json:"locale,omitempty"`
	AppVersion string `json:"appVersion,omitempty"`
}

// Device is a registered push recipient.
type Device struct {
	ID         string `json:"id"`
	AppID      string `json:"appId"`
	Token      string `json:"token"`
	Platform   string `json:"platform"`
	UserRef    string `json:"userRef,omitempty"`
	Locale     string `json:"locale,omitempty"`
	AppVersion string `json:"appVersion,omitempty"`
	LastSeenAt string `json:"lastSeenAt,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
}

// RegisterDevice registers a device token so it can receive pushes. Calling it
// again with the same token refreshes the device (upsert).
func (c *Client) RegisterDevice(ctx context.Context, appID string, req RegisterDeviceRequest) (*Device, error) {
	var d Device
	if err := c.do(ctx, http.MethodPost, "/v1/push/apps/"+esc(appID)+"/devices", nil, req, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// GetPushMessage returns a sent push's delivery summary (per-device outcomes).
// The response shape is returned as a decoded map to stay flexible.
func (c *Client) GetPushMessage(ctx context.Context, appID, messageID string) (map[string]any, error) {
	var out map[string]any
	if err := c.do(ctx, http.MethodGet, "/v1/push/apps/"+esc(appID)+"/messages/"+esc(messageID), nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}
