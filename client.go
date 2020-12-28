package nakivo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://%s:%d/c/router"

func NewClient(httpClient *http.Client, address string, port int) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	if address == "" {
		return nil, fmt.Errorf("unknown address")
	}
	if port <= 0 {
		return nil, fmt.Errorf("unknown port")
	}
	u, err := url.Parse(fmt.Sprintf(defaultBaseURL, address, port))
	if err != nil {
		return nil, err
	}
	client := &Client{
		baseURL:    u,
		httpClient: httpClient,
	}
	client.common.client = client
	client.Authentication = (*AuthenticationService)(&client.common)

	return client, nil
}

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client

	common         service
	Authentication *AuthenticationService
}

type service struct {
	client *Client
}

type request struct {
	// Action represents the action that is invoked.
	Action string `json:"action"`

	// Method represents the method that is invoked.
	Method string `json:"method"`

	// Type represents the type of the requests. Must be "rpc" at all times.
	Type string `json:"type"`

	// Tid represents the transactions id of a request. Used to identify the request by both the
	// client and the server. If the client sends a batch of requests, tid must be unique among
	// the requests
	Tid int `json:"tid"`

	// Data represents the parameters of the request. Format depends on the request type.
	Data interface{} `json:"data"`
}

type response struct {
	// Action represents the request action.
	Action string `json:"action"`

	// Method represents the requested method.
	Method string `json:"method"`

	// Tid represents the transactions id of a request.
	Tid string `json:"tid"`

	// Type represents the type of the requests. Must be "rpc" at all times.
	Type string `json:"type"`

	// Message if the request failed.
	Message string `json:"message,omitempty"`

	// Reference to the method where the problem occurred.
	Where string `json:"where,omitempty"`

	// Cause of failure.
	Cause string `json:"cause,omitempty"`

	Data interface{} `json:"data,omitempty"`
}

type APIError struct {
	Message string
	Where   string
	Cause   string
}

func (err *APIError) Error() string {
	if err.Where != "" {
		return fmt.Sprintf("api: request failed with %s on %s (%s)", err.Message, err.Where, err.Cause)
	}
	return fmt.Sprintf("api: request failed with '%s' (%s)", err.Message, err.Cause)
}

func (c *Client) NewRequest(request *request) (*http.Request, error) {
	r, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.baseURL.String(), bytes.NewReader(r))
	if err != nil {
		return nil, fmt.Errorf("new http request failed (%s)", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, response *response) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, fmt.Errorf("http client do failed (%s)", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		// todo: handle io.EOF on empty body
		return resp, err
	}
	if response.Message != "" && response.Cause != "" {
		return resp, &APIError{Message: response.Message, Where: response.Where, Cause: response.Cause}
	}
	return resp, nil
}
