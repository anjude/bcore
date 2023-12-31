package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient is a struct for configuring the HTTP client.
type HTTPClient struct {
	Client    *http.Client
	BaseURL   string
	Timeout   time.Duration
	UserAgent string
}

// NewHTTPClient creates a new HTTPClient with default configurations.
func NewHTTPClient(baseURL string, timeout time.Duration, userAgent string) *HTTPClient {
	return &HTTPClient{
		Client:    &http.Client{},
		BaseURL:   baseURL,
		Timeout:   timeout,
		UserAgent: userAgent,
	}
}

// SetTimeout sets the timeout for the HTTP client.
func (c *HTTPClient) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
	c.Client.Timeout = timeout
}

// SetUserAgent sets the User-Agent header for the HTTP client.
func (c *HTTPClient) SetUserAgent(userAgent string) {
	c.UserAgent = userAgent
}

// Request sends an HTTP request with the specified method, headers, request body, and response structure.
func (c *HTTPClient) Request(method, endpoint string, headers map[string]string, reqBody interface{}, respBody interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	// Prepare request body
	var body io.Reader
	if reqBody != nil {
		jsonPayload, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(jsonPayload)
	}

	// Create request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Set user-agent
	req.Header.Set("User-Agent", c.UserAgent)

	// Send request
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response body
	if respBody != nil {
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			return err
		}
	}

	return nil
}
