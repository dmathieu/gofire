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

		fmt.Fprint(w, `{"message": {"body": "hello", "type": "TextMessage"}}`)
	})

	res, err := client.Say("Some lambda message")
	assert.Equal(t, nil, err)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, `{"message": {"body": "hello", "type": "TextMessage"}}`, string(body))
}
