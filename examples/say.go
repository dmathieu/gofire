package main

import (
  "github.com/dmathieu/gofire/gofire"
)

func main() {

  client := gofire.Client{
    Token:     "<your API token>",
    Subdomain: "<your subdomain>",
    Room:    "<your room id (not it's name)>",
  }
  client.Say("Hello World!")
}
