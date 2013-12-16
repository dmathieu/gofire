package gofire

import (
	"bufio"
	"bytes"
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

func (c *Streaming) connect() (*bufio.Reader, error) {
	if c.stale {
		return nil, errors.New("Cannot connect on a stale client")
	}

	var tcpConn net.Conn
	tcpConn, err := net.Dial("tcp", c.path.Host+":443")
	if err != nil {
		return nil, err
	}
	ssl := tls.Client(tcpConn, nil)

	reader := bufio.NewReader(ssl)
	c.clientConn = httputil.NewClientConn(ssl, reader)

	req, err := http.NewRequest("GET", c.path.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", hashAuth(c.client.token, "x")))

	_, err = c.clientConn.Do(req)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

type listenCallback func([]byte)

func (c *Streaming) Listen(fun listenCallback) {
	reader, err := c.connect()
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
