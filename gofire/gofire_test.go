package gofire

import (
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server

	tlsMux    *http.ServeMux
	tlsCli    *Client
	tlsServer *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	tlsMux = http.NewServeMux()
	tlsServer = httptest.NewTLSServer(tlsMux)

	client = NewClientWith(server.URL, tlsServer.URL, "token")
}

// teardown closes the test HTTP server.
func tearDown() {
	server.Close()
	tlsServer.Close()
}
