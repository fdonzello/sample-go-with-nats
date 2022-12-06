package main

import (
	"fmt"
	"log"

	"go-with-nats.io/consumer/internal/messages"
)

func main() {
	nc, err := messages.Connect()
	if err != nil {
		log.Fatal(err)
	}

	var name string
	var surname string

	fmt.Print("user name: ")
	fmt.Scanf("%s", &name)

	fmt.Print("user lastname: ")
	fmt.Scanf("%s", &surname)

	if err := nc.Publish("users.create", fmt.Sprintf("%s %s", name, surname)); err != nil {
		log.Fatal(err)
	}

	if err := nc.Flush(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("user creation request sent!")
}
