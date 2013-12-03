package main

import (
	"github.com/dmathieu/gofire/gofire"
)

func main() {

	client := gofire.NewClient("<your API token>", "<your subdomain>", "<your room id (not it's name)>")
	client.Say("Hello World!")
}
