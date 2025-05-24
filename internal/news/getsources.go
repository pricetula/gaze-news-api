package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (n *News) GetSources() ([]Source, error) {
	// Construct the URL for the sources endpoint
	url := fmt.Sprintf("%s/top-headlines/sources?apiKey=%s", baseURL, n.apiKey)

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Execute the request using the client's http.Client
	resp, err := n.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request to NewsAPI: %w", err)
	}
	// Ensure the response body is closed
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode != http.StatusOK {
		// Attempt to unmarshal error response if available (NewsAPI returns JSON errors)
		var errorResp struct {
			Status  string `json:"status"`
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errorResp); err == nil {
			return nil, fmt.Errorf("NewsAPI error (status: %d, code: %s): %s", resp.StatusCode, errorResp.Code, errorResp.Message)
		}
		return nil, fmt.Errorf("NewsAPI request failed with status: %d, response: %s", resp.StatusCode, string(body))
	}

	// Unmarshal the JSON response into our SourcesResponse struct
	var sourcesResponse SourcesResponse
	if err := json.Unmarshal(body, &sourcesResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NewsAPI sources response: %w", err)
	}

	// Check the 'status' field in the JSON response
	if sourcesResponse.Status != "ok" {
		return nil, fmt.Errorf("NewsAPI response status not 'ok': %s", sourcesResponse.Status)
	}

	return sourcesResponse.Sources, nil
}
