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
	client = &Client{Token: "token", Subdomain: "gofire", Room: "1234", baseURL: server.URL}
}

// teardown closes the test HTTP server.
func tearDown() {
	server.Close()
}
