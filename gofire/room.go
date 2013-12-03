package gofire

import (
	"fmt"
	"net/http"
)

type Room struct {
	client  *Client
	room_id string
}

func (r *Room) getSayUrl() string {
	return r.client.baseURL + fmt.Sprintf("/room/%s/speak.json", r.room_id)
}

func (r *Room) Say(phrase string) (*http.Response, error) {
	message := Message{Type: "TextMesage", Body: phrase}
	request := Request{path: r.getSayUrl(), message: message, client: r.client}

	return request.Post()
}
