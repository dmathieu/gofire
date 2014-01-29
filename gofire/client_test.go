package gofire

import (
	"fmt"
	"github.com/bmizerany/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSuccessfulNew(t *testing.T) {
	client := NewClient("1234", "gofire", true)
	assert.Equal(t, "1234", client.token)
	assert.Equal(t, "https://gofire.campfirenow.com", client.baseURL)
}

func TestSuccessfulInsecureNew(t *testing.T) {
	client := NewClient("456", "gofire", false)
	assert.Equal(t, "456", client.token)
	assert.Equal(t, "http://gofire.campfirenow.com", client.baseURL)
}

func TestSuccessfulSearch(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "Basic dG9rZW46eA==", r.Header.Get("Authorization"))

		query := url.Values{"q": []string{"world"}, "format": []string{"json"}}
		assert.Equal(t, query, r.URL.Query())

		fmt.Fprint(w, `{"messages": [{"body": "hello", "type": "TextMessage"}]}`)
	})

	messages, err := client.Search("world")

	assert.Equal(t, nil, err)
	assert.Equal(t, "hello", messages[0].Body)
	assert.Equal(t, "TextMessage", messages[0].Type)
}

func TestSuccessfulRoomsList(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "Basic dG9rZW46eA==", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "Gofire (42@dmathieu.com)", r.Header.Get("User-Agent"))

		fmt.Fprint(w, `{"rooms":[{"id":1,"name":"gofire","topic":"an awesome campfire library for go"}]}`)
	})

	rooms, _ := client.Rooms()
	room := rooms[0]

	assert.Equal(t, len(rooms), 1)

	assert.Equal(t, 1, room.Id)
	assert.Equal(t, "gofire", room.Name)
	assert.Equal(t, "an awesome campfire library for go", room.Topic)
	assert.Equal(t, client, room.client)
}
