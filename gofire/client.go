package gofire

import (
	"fmt"
	"net/http"
)

var httpClient = &http.Client{}

type Client struct {
	http *http.Client

	token   string
	baseURL string
}

func NewClient(token, subdomain string, secure bool) *Client {
	var scheme string
	if secure {
		scheme = "https"
	} else {
		scheme = "http"
	}

	url := fmt.Sprintf(apiUrl, scheme, subdomain)

	return NewClientWith(url, token)
}

func NewClientWith(baseURL, token string) *Client {
	return &Client{token: token, baseURL: baseURL, http: httpClient}
}

func (c *Client) NewRoom(room_id string) *Room {
	return &Room{room_id: room_id, client: c}
}
