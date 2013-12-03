package gofire

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	path    string
	message Message
	client  *Client
}

func build_data(message Message) *strings.Reader {
	return strings.NewReader(message.Encode())
}

func hashAuth(user, password string) string {
	var a = fmt.Sprintf("%s:%s", user, password)
	return base64.StdEncoding.EncodeToString([]byte(a))
}

func (r *Request) Post() (*http.Response, error) {
	req, _ := http.NewRequest("POST", r.path, build_data(r.message))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", hashAuth(r.client.token, "x")))

	res, err := r.client.http.Do(req)
	return res, err
}
