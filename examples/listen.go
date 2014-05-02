package main

import (
	"fmt"
	"github.com/dmathieu/gofire/gofire"
)

func streaming() {
	client := gofire.NewClient("<your API token>", "<your subdomain>", true)
	room := client.NewRoom(15) //15 being your room id (not it's name)>

	channel, err := room.Listen()
	if err != nil {
		fmt.Println("Listening err:", err)
		return
	}
	for {
		msg := <-channel
		fmt.Println(msg.Body)
	}
}
