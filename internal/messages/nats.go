package messages

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

type Message interface {
}

func Connect() (*nats.EncodedConn, error) {

	fmt.Println("connecting to NATS host", os.Getenv("NATS_HOST"))

	nc, err := nats.Connect(fmt.Sprintf("nats://%s:4222", os.Getenv("NATS_HOST")))
	if err != nil {
		return nil, err
	}

	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	// Do something with the connection
	return ec, nil
}

func ReceiveMessages[K Message](URL string, subject string, applyFunc func(msg K) error) chan error {
	errCh := make(chan error)

	nc, err := Connect()
	if err != nil {
		errCh <- err
		return errCh
	}

	ch := make(chan K)
	nc.BindRecvChan(subject, ch)

	for {
		select {
		case msg := <-ch:
			if err := applyFunc(msg); err != nil {
				errCh <- err
				continue
			}
		}
	}

}
