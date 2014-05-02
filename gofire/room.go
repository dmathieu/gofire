package gofire

import (
	"log"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
)

type Room struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Topic string `json:"topic"`

	client *Client
}

func (r *Room) getSayUrl() string {
	return fmt.Sprintf("%s/room/%d/speak.json", r.client.baseURL, r.Id)
}

func (r *Room) getStreamUrl() (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s/room/%d/live.json", r.client.streamingBaseURL, r.Id))
}

func (r *Room) Say(phrase string) (*Message, error) {
	subject := map[string]Message{
		"message": Message{Body: phrase},
	}
	request := Request{path: r.getSayUrl(), subject: subject, client: r.client}
	response, err := request.Post()
	if err != nil {
		return nil, err
	}

	var jsonRoot map[string]Message
	err = response.UnmarshalJSON(&jsonRoot)
	if err != nil {
		return nil, err
	}

	message := jsonRoot["message"]
	return &message, nil
}

func (r *Room) startListener(channel chan Message) {
	url, err := r.getStreamUrl()
	if err != nil {
		log.Fatal("Room.startListener():", err)
	}
	stream := Streaming{path: url, dial: net.Dial, client: r.client}

	stream.Listen(func(content []byte) {
		var message Message

		err := json.Unmarshal(content, &message)
		if err == nil {
			channel <- message
		}
	})
}

func (r *Room) Listen() chan Message {
	channel := make(chan Message)
	go r.startListener(channel)

	return channel
}
