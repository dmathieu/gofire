package gofire

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	http *http.Response
}

func (r *Response) ReadBody() []byte {
	defer r.http.Body.Close()
	body, err := ioutil.ReadAll(r.http.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func (r *Response) UnmarshalJSON(subject interface{}) error {
	body := r.ReadBody()
	return json.Unmarshal(body, subject)
}
