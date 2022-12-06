package main

import "github.com/nats-io/nats.go"

func Connect(URL string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(URL)
	if err != nil {
		return nil, err
	}

	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	// Do something with the connection
	return ec, nil
}

func main() {

}
