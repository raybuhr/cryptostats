package cryptostats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIClient helps build connections to cryptocurrency APIs
type APIClient struct {
	Client *http.Client
	Token  string
}

// New instantiates a http connection for the Client
func New(token string) *APIClient {
	return &APIClient{
		Client: &http.Client{},
		Token:  token,
	}
}

// Do sends an HTTP request and returns an HTTP response for the Client
func (ac *APIClient) Do(method, url string, bodyObj, responseObj interface{}) error {
	var body io.Reader // nil until it is assigned
	if body != nil {
		b, err := json.Marshal(bodyObj)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Use Client.Do from net/http to make request
	resp, err := ac.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("API Returned HTTP Status Code of %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(responseObj); err != nil {
		return err
	}
	return nil
}
