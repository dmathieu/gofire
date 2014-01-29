package gofire

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
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

func (r *Request) Do(verb string) (*Response, error) {
	marshalled, err := json.Marshal(r.subject)
	if err != nil {
		panic(err)
	}
	content := bytes.NewReader(marshalled)

	req, err := http.NewRequest(verb, r.path, content)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", hashAuth(r.client.token, "x")))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	res, err := r.client.http.Do(req)

	if err == nil && res.StatusCode < 200 || res.StatusCode > 209 {
		err = errors.New(fmt.Sprintf("Invalid status code: %s", res.Status))
	}

	response := &Response{http: res}
	return response, err
}

func (r *Request) Post() (*Response, error) {
	return r.Do("POST")
}

func (r *Request) Get() (*Response, error) {
	return r.Do("GET")
}
