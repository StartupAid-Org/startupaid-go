package startupaid

import (
	"context"
	"net/http"
)

// SendOTPRequest is the payload for SendOTP. Only To is required — we generate
// and deliver the code and track verification for you. Supply Code instead to
// deliver a code you generated yourself (delivery-only mode).
type SendOTPRequest struct {
	To      string `json:"to"`                // recipient phone in E.164, e.g. "+2348012345678"
	AppName string `json:"appName,omitempty"` // shown in the message; defaults to your OTP settings
	Code    string `json:"code,omitempty"`    // deliver this code as-is instead of generating one
	Region  string `json:"region,omitempty"`  // sender region override, e.g. "NG"; omit to auto-route
}

// OTPChallenge is returned by SendOTP.
type OTPChallenge struct {
	ID        string `json:"id"`
	To        string `json:"to"`
	Status    string `json:"status"`    // pending | verified | expired | failed | canceled
	ExpiresAt string `json:"expiresAt"` // RFC 3339
}

// SendOTP generates and delivers a one-time passcode over WhatsApp. Metered as
// one unit per send.
func (c *Client) SendOTP(ctx context.Context, req SendOTPRequest) (*OTPChallenge, error) {
	var ch OTPChallenge
	if err := c.do(ctx, http.MethodPost, "/v1/otp/send", nil, req, &ch); err != nil {
		return nil, err
	}
	return &ch, nil
}

// VerifyOTP checks a code the recipient submitted for a phone number. Returns
// true when it matches an active challenge. Verification is free (not metered).
func (c *Client) VerifyOTP(ctx context.Context, to, code string) (bool, error) {
	var out struct {
		Verified bool `json:"verified"`
	}
	body := map[string]string{"to": to, "code": code}
	if err := c.do(ctx, http.MethodPost, "/v1/otp/verify", nil, body, &out); err != nil {
		return false, err
	}
	return out.Verified, nil
}
