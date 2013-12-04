package gofire

import (
	"encoding/json"
	"fmt"
)

type Room struct {
	client  *Client
	room_id string
}

func (r *Room) getSayUrl() string {
	return r.client.baseURL + fmt.Sprintf("/room/%s/speak.json", r.room_id)
}

func (r *Room) Say(phrase string) (*Message, error) {
	message := Message{Type: "TextMessage", Body: phrase}
	request := Request{path: r.getSayUrl(), subject: message, client: r.client}
	response, err := request.Post()
	if err != nil {
		panic(err)
	}

	body := response.ReadBody()
	err = json.Unmarshal(body, &message)

	return &message, err
}
