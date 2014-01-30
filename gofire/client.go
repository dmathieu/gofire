package gofire

import (
	"fmt"
	"net/http"
)

var httpClient = &http.Client{}

type Client struct {
	http *http.Client

	token            string
	baseURL          string
	streamingBaseURL string
}

func NewClient(token, subdomain string, secure bool) *Client {
	var scheme string
	if secure {
		scheme = "https"
	} else {
		scheme = "http"
	}

	url := fmt.Sprintf(apiUrl, scheme, subdomain)

	return NewClientWith(url, streamingBaseURL, token)
}

func NewClientWith(baseURL, streamingBaseURL, token string) *Client {
	return &Client{token: token, baseURL: baseURL, streamingBaseURL: streamingBaseURL, http: httpClient}
}

func (c *Client) NewRoom(room_id int) *Room {
	return &Room{Id: room_id, client: c}
}

func (c *Client) getSearchUrl(query string) string {
	return fmt.Sprintf("%s/search?q=%s&format=json", c.baseURL, query)
}

func (c *Client) getRoomsUrl() string {
	return fmt.Sprintf("%s/rooms", c.baseURL)
}

func (c *Client) Search(query string) ([]Message, error) {
	request := Request{path: c.getSearchUrl(query), subject: nil, client: c}
	response, err := request.Get()
	if err != nil {
		return nil, err
	}

	var jsonRoot map[string][]Message
	err = response.UnmarshalJSON(&jsonRoot)
	if err != nil {
		return nil, err
	}

	return jsonRoot["messages"], err
}

func (c *Client) Rooms() ([]Room, error) {
	request := Request{path: c.getRoomsUrl(), subject: nil, client: c}
	response, err := request.Get()
	if err != nil {
		return nil, err
	}

	var jsonRoot map[string][]Room
	err = response.UnmarshalJSON(&jsonRoot)
	if err != nil {
		return nil, err
	}

	rooms := jsonRoot["rooms"]
	for index := range rooms {
		rooms[index].client = c
	}

	return rooms, err
}
