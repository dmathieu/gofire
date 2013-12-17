package gofire

import (
	"encoding/json"
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

func (c *Client) getSearchUrl(query string) string {
	return fmt.Sprintf("%s/search?q=%s&format=json", c.baseURL, query)
}

func (c *Client) Search(query string) ([]Message, error) {
	request := Request{path: c.getSearchUrl(query), subject: nil, client: c}
	response, err := request.Get()
	if err != nil {
		panic(err)
	}

	var jsonRoot map[string][]Message
	body := response.ReadBody()

	err = json.Unmarshal(body, &jsonRoot)
	if err != nil {
		panic(err)
	}

	return jsonRoot["messages"], err
}
