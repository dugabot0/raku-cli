package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	rakutenBaseURL      = "https://app.rakuten.co.jp/services/api/"
	ichibaOpenAPIURL    = "https://openapi.rakuten.co.jp/ichibams/api/"
)

// Client handles requests to the Rakuten Web Service API.
type Client struct {
	AppID       string
	AffiliateID string
	AccessKey   string // optional; enables Ichiba OpenAPI endpoint
	Origin      string // optional; sent as Origin header (required by some app registrations)
	HTTP        *http.Client
}

// New constructs a Client with the given timeout.
func New(appID, affiliateID, accessKey, origin string, timeout time.Duration) *Client {
	return &Client{
		AppID:       appID,
		AffiliateID: affiliateID,
		AccessKey:   accessKey,
		Origin:      origin,
		HTTP:        &http.Client{Timeout: timeout},
	}
}

// ichibaGet routes to the OpenAPI endpoint if AccessKey is set,
// otherwise falls back to the standard Rakuten API.
func (c *Client) ichibaGet(standardEndpoint, openAPIEndpoint string, params url.Values) (json.RawMessage, error) {
	if c.AccessKey != "" {
		params.Set("accessKey", c.AccessKey)
		u := ichibaOpenAPIURL + openAPIEndpoint + "?" + params.Encode()
		return c.fetch(u)
	}
	return c.get(standardEndpoint, params)
}

func (c *Client) baseParams() url.Values {
	p := url.Values{}
	p.Set("applicationId", c.AppID)
	p.Set("format", "json")
	p.Set("formatVersion", "2")
	if c.AffiliateID != "" {
		p.Set("affiliateId", c.AffiliateID)
	}
	return p
}

func (c *Client) get(endpoint string, params url.Values) (json.RawMessage, error) {
	return c.fetch(rakutenBaseURL + endpoint + "?" + params.Encode())
}

func (c *Client) fetch(u string) (json.RawMessage, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, &APIError{StatusCode: 0, Message: fmt.Sprintf("build request: %v", err)}
	}
	if c.Origin != "" {
		req.Header.Set("Origin", c.Origin)
		req.Header.Set("Referer", c.Origin+"/")
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, &APIError{StatusCode: 0, Message: fmt.Sprintf("network error: %v", err)}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// ok
	case http.StatusUnauthorized, http.StatusForbidden:
		return nil, &AuthError{StatusCode: resp.StatusCode}
	default:
		return nil, &APIError{StatusCode: resp.StatusCode}
	}

	var raw json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return raw, nil
}

// AuthError is returned on HTTP 401/403.
type AuthError struct{ StatusCode int }

func (e *AuthError) Error() string {
	return fmt.Sprintf("authentication error (HTTP %d) — check RAKUTEN_APP_ID", e.StatusCode)
}

// APIError is returned on non-OK HTTP responses.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("API error (HTTP %d)", e.StatusCode)
}

// setIfNonEmpty sets key=val in p if val is non-empty.
func setIfNonEmpty(p url.Values, key, val string) {
	if val != "" {
		p.Set(key, val)
	}
}
