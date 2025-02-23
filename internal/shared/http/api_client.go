package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const timeout = 5 * time.Second

type ApiClient struct {
	client   http.Client
	endpoint string
}

// NewApiClient creates a new `ApiClient` instance.
func NewApiClient(endpoint string) *ApiClient {

	client := http.Client{}
	return &ApiClient{
		client,
		endpoint,
	}
}

// Get makes a GET request to the API endpoint, on the provided path.
func (apic *ApiClient) Get(path string) (respBody []byte, err error) {
	return apic.doRequest(context.Background(), http.MethodGet, path, nil, nil)
}

func (apic *ApiClient) doRequest(ctx context.Context, method, path string, params map[string]string, reqBody []byte) (respBody []byte, err error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	u, err := url.Parse(apic.endpoint + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request URL: %w", err)
	}
	// Encode the query params.
	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// Perform the HTTP request.
	resp, err := apic.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do the HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP response body: %w", err)
	}

	// check if status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with status code: %d: %w",
			resp.StatusCode, errors.New(string(respBody)))
	}

	return respBody, nil
}
