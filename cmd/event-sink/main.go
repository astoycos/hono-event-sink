package main

import (
	"context"
	"log"
	"pack.ag/amqp"
)

func main() {
	client, err := amqp.Dial("amqps://localhost:5671", amqp.ConnSASLPlain("test", "test"))
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()

	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress("/events"),
		amqp.LinkCredit(10))
	if err != nil {
		log.Fatal("Creating AMQP receiver:", err)
	}

	msg, err := receiver.Receive(ctx)
	if err != nil {
		log.Fatal("Reading message:", err)
	}
	// TODO: Store in sqlite3 db

	msg.Accept()
}