package gofire

import (
	"testing"
	"net/http"
	"fmt"
)

var httpErr = fmt.Errorf("an error")

type MockClient struct {
}

func (c MockClient) Do(req *http.Request) (*http.Response, error) {
	return nil, httpErr
}

// Make sure we're handling errors in the http client
func TestDo_HttpErrsHandled(t *testing.T) {
	c := &Client{
		http: MockClient{},
	}

	req := Request{path: "", subject: map[string]Message{}, client: c}

	_, err := req.Do("something")

	if err != httpErr {
		t.Fatal("Expected %v; got %v", httpErr, err)
	}
}
