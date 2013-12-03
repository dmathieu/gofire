package gofire

import (
	"fmt"
	"net/http"
)

type Client struct {
	Token, Subdomain, Room string
	baseURL                string
}

func (c *Client) getBaseURL() string {
	if c.baseURL == "" {
		c.baseURL = fmt.Sprintf(apiUrl, c.Subdomain)
	}

	return c.baseURL
}

func (c *Client) getSayUrl() string {
	return c.getBaseURL() + fmt.Sprintf("/room/%s/speak.json", c.Room)
}

func (c *Client) Say(phrase string) (*http.Response, error) {
	path := c.getSayUrl()
	message := Message{Type: "TextMesage", Body: phrase}

	request := Request{Path: path, Message: message, Client: c}

	return request.Post()
}
