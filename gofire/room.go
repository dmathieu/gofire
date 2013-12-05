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

func (r *Room) Say(phrase string) (Message, error) {
	subject := map[string]Message{
		"message": Message{Type: "TextMessage", Body: phrase},
	}
	request := Request{path: r.getSayUrl(), subject: subject, client: r.client}
	response, err := request.Post()
	if err != nil {
		panic(err)
	}

	var jsonRoot map[string]Message
	body := response.ReadBody()
	err = json.Unmarshal(body, &jsonRoot)

	return jsonRoot["message"], err
}
