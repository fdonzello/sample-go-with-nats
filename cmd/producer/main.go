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
	var lastname string

	fmt.Print("user name: ")
	fmt.Scanf("%s", &name)

	fmt.Print("user lastname: ")
	fmt.Scanf("%s", &lastname)

	if len(name) == 0 {
		log.Fatal("empty name provided")
	}

	if len(lastname) == 0 {
		log.Fatal("empty lastname provided")
	}

	if err := nc.Publish("users.create", fmt.Sprintf("%s %s", name, lastname)); err != nil {
		log.Fatal(err)
	}

	if err := nc.Flush(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("user creation request sent!")
}
