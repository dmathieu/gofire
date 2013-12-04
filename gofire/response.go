package gofire

import (
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
