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
	"strings"
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

func hasPort(host string) bool {
	parts := strings.Split(host, ":")
	// make sure there is 1 colon, and that either side of the colon is non-empty
	return len(parts) == 2 && parts[0] != "" && parts[1] != ""
}

func (c *Streaming) connect() (*bufio.Reader, error) {
	if c.stale {
		return nil, errors.New("Cannot connect on a stale client")
	}

	var tcpConn net.Conn
	host := c.path.Host
	if !hasPort(host) {
		host = host + ":443"
	}

	tcpConn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	config := &tls.Config{InsecureSkipVerify: true}
	ssl := tls.Client(tcpConn, config)

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

func (c *Streaming) Listen(fun listenCallback) error {
	reader, err := c.connect()
	if err != nil {
		return err
	}

	for {
		if c.stale {
			c.clientConn.Close()
			break
		}

		line, err := reader.ReadBytes('\r')
		if err != nil {
			c.stale = true
		} else {
			line = bytes.TrimSpace(line)

			fun(line)
		}
	}

	return nil
}
