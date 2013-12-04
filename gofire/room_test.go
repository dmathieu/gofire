package gofire

import (
	"fmt"
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestSuccessfulSay(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/room/1234/speak.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Basic dG9rZW46eA==", r.Header.Get("Authorization"))

		fmt.Fprint(w, `{"message": {"body": "hello", "type": "TextMessage"}}`)
	})

	room := client.NewRoom("1234")
	message, err := room.Say("Some lambda message")

	assert.Equal(t, nil, err)
	assert.Equal(t, "Some lambda message", message.Body)
	assert.Equal(t, "TextMessage", message.Type)
}
