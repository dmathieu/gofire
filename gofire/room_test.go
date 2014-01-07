package gofire

import (
	"fmt"
	"github.com/bmizerany/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSuccessfulSay(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/room/1234/speak.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Basic dG9rZW46eA==", r.Header.Get("Authorization"))
		assert.Equal(t, "Gofire (42@dmathieu.com)", r.Header.Get("User-Agent"))

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, `{"message":{"type":"TextMessage","body":"Some lambda message"}}`, string(body))

		fmt.Fprint(w, `{"message": {"body": "hello", "type": "TextMessage"}}`)
	})

	room := client.NewRoom("1234")
	message, err := room.Say("Some lambda message")

	assert.Equal(t, nil, err)
	assert.Equal(t, "hello", message.Body)
	assert.Equal(t, "TextMessage", message.Type)
}

func TestSuccessfulListen(t *testing.T) {
	setup()
	defer tearDown()

	tlsMux.HandleFunc("/room/1234/live.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		fmt.Fprint(w, `{"body": "hello", "type": "TextMessage"}`)
		fmt.Fprint(w, "\r")
		fmt.Fprint(w, `{"body": "world", "type": "TextMessage"}`)
		fmt.Fprint(w, "\r")
	})

	room := client.NewRoom("1234")
	channel := room.Listen()

	msg := <-channel
	assert.Equal(t, "hello", msg.Body)
	assert.Equal(t, "TextMessage", msg.Type)

	msg = <-channel
	assert.Equal(t, "world", msg.Body)
	assert.Equal(t, "TextMessage", msg.Type)
}
