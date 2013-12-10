package main

import (
	"fmt"
	"github.com/dmathieu/gofire/gofire"
)

func streaming() {

	client := gofire.NewClient("<your API token>", "<your subdomain>", true)
	room := client.NewRoom("<your room id (not it's name)>")

	channel := room.Listen()
	for {
		msg := <-channel
		fmt.Println(msg.Body)
	}
}
