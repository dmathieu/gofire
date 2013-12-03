package gofire

import (
  "net/url"
)

type Message struct {
  Type, Body string
}

func (m *Message) Encode() (string) {
  data := url.Values{"message[type]": {m.Type}, "message[body]": {m.Body}}
  return data.Encode()
}
