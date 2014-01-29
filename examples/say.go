package main

import (
	"fmt"
	"github.com/dmathieu/gofire/gofire"
)

func say() {

	client := gofire.NewClient("<your API token>", "<your subdomain>", true)

	rooms, _ := client.Rooms()
	// Or retrieve a single room by it's id
	//room := client.NewRoom(<your room id (not it's name)>)

	for _, room := range rooms {
		room.Say(fmt.Sprintf("Hello room %s!", room.Name))
	}
}
