package startupaid

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// ConversionResult is the outcome of a Convert call.
type ConversionResult struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Amount   float64 `json:"amount"`
	Result   float64 `json:"result"`
	Rate     float64 `json:"rate"`
	RateAsOf string  `json:"rateAsOf"`
	// Set when a custom account rate/spread applied.
	MarketRate   float64  `json:"marketRate,omitempty"`
	MarketResult float64  `json:"marketResult,omitempty"`
	Spread       *float64 `json:"spread,omitempty"`
}

// Convert converts amount from one currency to another at live rates.
func (c *Client) Convert(ctx context.Context, from, to string, amount float64) (*ConversionResult, error) {
	q := url.Values{}
	q.Set("from", from)
	q.Set("to", to)
	q.Set("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	var r ConversionResult
	if err := c.do(ctx, http.MethodGet, "/v1/convert", q, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// Rates returns exchange rates for a base currency against all supported
// currencies. The response is returned as a decoded map to stay flexible.
func (c *Client) Rates(ctx context.Context, base string) (map[string]any, error) {
	var out map[string]any
	if err := c.do(ctx, http.MethodGet, "/v1/rates/"+esc(base), nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Currencies lists the supported currencies.
func (c *Client) Currencies(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	if err := c.do(ctx, http.MethodGet, "/v1/currencies", nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}
