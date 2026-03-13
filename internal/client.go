package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultBaseURL = "https://app.launchdarkly.com/api/v2"

// launchDarklyClient wraps an HTTP client for the LaunchDarkly REST API v2.
type launchDarklyClient struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// newLaunchDarklyHTTPClient creates a new launchDarklyClient.
func newLaunchDarklyHTTPClient(apiKey, baseURL string) *launchDarklyClient {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &launchDarklyClient{
		httpClient: &http.Client{},
		baseURL:    baseURL,
		apiKey:     apiKey,
	}
}

// doRequest executes an HTTP request against the LaunchDarkly API.
// body may be nil for requests with no payload.
func (c *launchDarklyClient) doRequest(ctx context.Context, method, path string, body any) (map[string]any, error) {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("launchdarkly: marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("launchdarkly: create request: %w", err)
	}
	req.Header.Set("Authorization", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("LD-API-Version", "20220603")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("launchdarkly: do request: %w", err)
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("launchdarkly: read response: %w", err)
	}

	if resp.StatusCode == http.StatusNoContent || len(respData) == 0 {
		return map[string]any{"status": resp.StatusCode}, nil
	}

	if resp.StatusCode >= 400 {
		var errBody map[string]any
		if jsonErr := json.Unmarshal(respData, &errBody); jsonErr == nil {
			if msg, ok := errBody["message"].(string); ok {
				return nil, fmt.Errorf("launchdarkly API error %d: %s", resp.StatusCode, msg)
			}
		}
		return nil, fmt.Errorf("launchdarkly API error %d: %s", resp.StatusCode, string(respData))
	}

	var result map[string]any
	if err := json.Unmarshal(respData, &result); err != nil {
		// May be an array at top level
		var arr []any
		if err2 := json.Unmarshal(respData, &arr); err2 == nil {
			return map[string]any{"items": arr, "count": len(arr)}, nil
		}
		return nil, fmt.Errorf("launchdarkly: parse response: %w", err)
	}
	return result, nil
}

// doRequestList executes an HTTP request and returns a list-style response.
// LaunchDarkly list endpoints return {"items": [...], "totalCount": N, "_links": {...}}.
func (c *launchDarklyClient) doRequestList(ctx context.Context, path string) (map[string]any, error) {
	return c.doRequest(ctx, http.MethodGet, path, nil)
}
