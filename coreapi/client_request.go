package coreapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Joins provided url parts with the base url to build the endpoint for calling
func (c *Client) URL(parts ...string) string {
	path, _ := url.JoinPath(c.apiUrl, parts...)
	return path
}

func (c *Client) NewRequest(
	ctx context.Context,
	method,
	uri string,
	payload any,
	options ...RequestOption,
) (*http.Request, error) {
	var body bytes.Buffer
	if payload != nil {
		if err := json.NewEncoder(&body).Encode(payload); err != nil {
			return nil, fmt.Errorf("encoding request payload failed: %w", err)
		}
	}

	request, err := http.NewRequestWithContext(ctx, method, uri, &body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	for _, option := range options {
		option(request)
	}

	return request, nil
}

// Do triggers an HTTP request and returns an HTTP response,
// handling any context cancellations or timeouts.
func (m *Client) Do(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	response, err := m.http.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	return response, nil
}

// Request combines NewRequest and Do, while also handling decoding of response payload.
//
// Can be used to provide a custom payload type for a request
func (m *Client) Request(ctx context.Context, method, uri string, payload any, result any, options ...RequestOption) error {
	request, err := m.NewRequest(ctx, method, uri, payload, options...)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}

	response, err := m.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}
	defer response.Body.Close()

	// If the response contains a client or a server error then return the error.
	if response.StatusCode >= http.StatusBadRequest {
		return newError(response)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	if result != nil {
		if err = json.Unmarshal(responseBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response payload: %w", err)
		}
	}

	return nil
}
