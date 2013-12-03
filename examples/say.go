package main

import (
	"github.com/dmathieu/gofire/gofire"
)

func main() {

	client := gofire.NewClient("<your API token>", "<your subdomain>", true)
	room := client.NewRoom("<your room id (not it's name)>")
	room.Say("Hello World!")
}
