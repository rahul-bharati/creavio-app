package client

import (
	"io"
	"net/http"
)

type NotificationClient struct {
	BaseURL string
	HTTP    *http.Client
}

func (c *NotificationClient) GetGreeting() (string, error) {
	resp, err := c.HTTP.Get(c.BaseURL + "/hello")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	greeting := string(body)
	return greeting, nil
}
