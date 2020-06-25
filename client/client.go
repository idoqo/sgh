package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	BaseUrl = "https://api.swaggerhub.com/"
)
// Client wraps a http client to talk to the SwaggerHub API
type Client struct {
	http    *http.Client
	baseUrl string
	authToken string
}

func NewClient(authToken string) *Client {
	tr := http.DefaultTransport

	return &Client{
		http: &http.Client{Transport: tr},
		baseUrl:   BaseUrl,
		authToken: authToken,
	}
}

// MakeRequest performs a HTTP request and parses the response
func (c Client) MakeRequest(method string, path string, body io.Reader, data interface{}) error {
	url := c.baseUrl + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	ok := res.StatusCode >= 200 && res.StatusCode < 300
	if !ok {
		return httpError(res)
	}

	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	return nil
}

// SetAuthToken sets the authentication token to be used in client requests
func (c *Client) SetAuthToken(token string) {
	c.authToken = token
}

func (c Client) AuthToken() string {
	return c.authToken
}

func httpError(res *http.Response) error {
	var message string
	var parsed struct {
		Message string `json:"message"`
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &parsed)
	if err != nil {
		message = string(b)
	} else {
		message = parsed.Message
	}

	return fmt.Errorf("http error, '%s' failed (%d): '%s'", res.Request.URL, res.StatusCode, message)
}
