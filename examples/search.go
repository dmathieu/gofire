package main

import (
	"fmt"
	"github.com/dmathieu/gofire/gofire"
)

func search() {

	client := gofire.NewClient("<your API token>", "<your subdomain>", true)
	messages, _ := client.Search("Search Query")

	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i].Body)
	}
}
