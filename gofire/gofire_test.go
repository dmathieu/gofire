package gofire

import (
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClientWith(server.URL, "token", "1234")
}

// teardown closes the test HTTP server.
func tearDown() {
	server.Close()
}
