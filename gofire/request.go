package gofire

import (
  "net/http"
  "encoding/base64"
  "strings"
  "fmt"
)

type Request struct {
  Path    string
  Message Message
  Client  *Client
}

func build_data(message Message) (*strings.Reader) {
  return strings.NewReader(message.Encode())
}

func hashAuth(user, password string) string {
  var a = fmt.Sprintf("%s:%s", user, password)
  return base64.StdEncoding.EncodeToString([]byte(a))
}

func (r *Request) Post() (*http.Response, error) {
  client   := &http.Client{}
  req, _ := http.NewRequest("POST", r.Path, build_data(r.Message))
  req.Header.Set("Authorization", fmt.Sprintf("Basic %s", hashAuth(r.Client.Token, "x")))

  res, err := client.Do(req)
  return res, err
}
