package gofire

import (
	"fmt"
	"net/http"
)

type Client struct {
	Token   string
	Room    string
	baseURL string
}

func NewClient(token, subdomain, room string) *Client {
	url := fmt.Sprintf(apiUrl, subdomain)

	return NewClientWith(url, token, room)
}

func NewClientWith(baseURL, token, room string) *Client {
	return &Client{Token: token, Room: room, baseURL: baseURL}
}

func (c *Client) getSayUrl() string {
	return c.baseURL + fmt.Sprintf("/room/%s/speak.json", c.Room)
}

func (c *Client) Say(phrase string) (*http.Response, error) {
	path := c.getSayUrl()
	message := Message{Type: "TextMesage", Body: phrase}

	request := Request{Path: path, Message: message, Client: c}

	return request.Post()
}
