package gofire

import (
  "fmt"
)

type Client struct {
  Token, Subdomain, Room string
}

func (c *Client) Say(phrase string) (bool) {
  path := fmt.Sprintf(apiUrl, c.Subdomain) + fmt.Sprintf("/room/%s/speak.json", c.Room)
  message := Message{Type: "TextMesage", Body: phrase}

  request := Request{Path: path, Message: message, Client: c}

  _, err := request.Post()
  return err == nil
}
