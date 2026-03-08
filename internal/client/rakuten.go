package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Each API group has its own base URL on openapi.rakuten.co.jp
const (
	ichibaItemsBaseURL   = "https://openapi.rakuten.co.jp/ichibams/api/"
	ichibaGenreBaseURL   = "https://openapi.rakuten.co.jp/ichibagt/api/"
	ichibaRankingBaseURL = "https://openapi.rakuten.co.jp/ichibaranking/api/"
	servicesBaseURL      = "https://openapi.rakuten.co.jp/services/api/"
	engineBaseURL        = "https://openapi.rakuten.co.jp/engine/api/"
	recipeBaseURL        = "https://openapi.rakuten.co.jp/recipems/api/"
)

// Client handles requests to the Rakuten Web Service API.
type Client struct {
	AppID       string
	AffiliateID string
	AccessKey   string
	Origin      string
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

func (c *Client) baseParams() url.Values {
	p := url.Values{}
	p.Set("applicationId", c.AppID)
	p.Set("format", "json")
	p.Set("formatVersion", "2")
	if c.AffiliateID != "" {
		p.Set("affiliateId", c.AffiliateID)
	}
	if c.AccessKey != "" {
		p.Set("accessKey", c.AccessKey)
	}
	return p
}

func (c *Client) ichibaGet(endpoint string, params url.Values) (json.RawMessage, error) {
	return c.fetch(ichibaItemsBaseURL + endpoint + "?" + params.Encode())
}

func (c *Client) ichibaGenreGet(endpoint string, params url.Values) (json.RawMessage, error) {
	return c.fetch(ichibaGenreBaseURL + endpoint + "?" + params.Encode())
}

func (c *Client) ichibaRankingGet(endpoint string, params url.Values) (json.RawMessage, error) {
	return c.fetch(ichibaRankingBaseURL + endpoint + "?" + params.Encode())
}

// get uses services base URL (Books, Kobo, BooksGenre, ...)
func (c *Client) get(endpoint string, params url.Values) (json.RawMessage, error) {
	return c.fetch(servicesBaseURL + endpoint + "?" + params.Encode())
}

// engineGet uses engine base URL (Travel, GORA)
func (c *Client) engineGet(endpoint string, params url.Values) (json.RawMessage, error) {
	return c.fetch(engineBaseURL + endpoint + "?" + params.Encode())
}

// recipeGet uses recipe base URL
func (c *Client) recipeGet(endpoint string, params url.Values) (json.RawMessage, error) {
	return c.fetch(recipeBaseURL + endpoint + "?" + params.Encode())
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
	return fmt.Sprintf("authentication error (HTTP %d) — check RAKUTEN_APP_ID / RAKUTEN_ACCESS_KEY", e.StatusCode)
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
