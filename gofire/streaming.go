package gofire

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Streaming struct {
	client     *Client
	path       *url.URL
	clientConn *httputil.ClientConn
	stale      bool
}

func (c *Streaming) Close() {
	c.stale = true
}

func (c *Streaming) connect() (*http.Response, error) {
	if c.stale {
		return nil, errors.New("Cannot connect on a stale client")
	}

	var tcpConn net.Conn
	tcpConn, err := net.Dial("tcp", c.path.Host+":443")
	if err != nil {
		return nil, err
	}
	cf := &tls.Config{Rand: rand.Reader}
	ssl := tls.Client(tcpConn, cf)

	c.clientConn = httputil.NewClientConn(ssl, nil)

	req, err := http.NewRequest("GET", c.path.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", hashAuth(c.client.token, "x")))

	resp, err := c.clientConn.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type listenCallback func([]byte)

func (c *Streaming) Listen(fun listenCallback) {
	resp, err := c.connect()
	reader := bufio.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}

	for {
		if c.stale {
			c.clientConn.Close()
			break
		}

		line, err := reader.ReadBytes('\r')
		if err != nil {
			panic(err)
		}
		line = bytes.TrimSpace(line)

		fun(line)
	}
}
