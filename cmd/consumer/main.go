package main

import (
	"fmt"
	"log"
	"os"

	"go-with-nats.io/consumer/internal/messages"
)

func main() {
	nc, err := messages.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Getenv("OUTPUT_FILE_PATH")) == 0 {
		log.Fatal("please set the OUTPUT_FILE_PATH env var")
	}

	f, err := os.OpenFile(os.Getenv("OUTPUT_FILE_PATH"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error opening the output file: %v", err)
	}

	fmt.Println("connected to nats")

	ch := make(chan string)

	if _, err := nc.BindRecvChan("users.create", ch); err != nil {
		log.Fatal(err)
	}

	fmt.Println("listening for messages")
	for {
		select {
		case a := <-ch:
			log.Println("received a user to create: ", a)

			if _, err := f.WriteString(a); err != nil {
				log.Fatal(err)
			}

			if _, err := f.WriteString("\n"); err != nil {
				log.Fatal(err)
			}
			log.Println("User created!")
		}
	}

}
