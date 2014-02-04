package gofire

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	http *http.Response
}

func (r *Response) ReadBody() ([]byte, error) {
	defer r.http.Body.Close()
	body, err := ioutil.ReadAll(r.http.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *Response) UnmarshalJSON(subject interface{}) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, subject)
}
