package gofire

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	path    string
	subject interface{}
	client  *Client
}

func hashAuth(user, password string) string {
	var a = fmt.Sprintf("%s:%s", user, password)
	return base64.StdEncoding.EncodeToString([]byte(a))
}

func (r *Request) Post() (*Response, error) {
	marshalled, err := json.Marshal(r.subject)
	if err != nil {
		panic(err)
	}
	content := bytes.NewReader(marshalled)

	req, _ := http.NewRequest("POST", r.path, content)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", hashAuth(r.client.token, "x")))

	res, err := r.client.http.Do(req)
	response := &Response{http: res}
	return response, err
}
