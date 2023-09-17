package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Test() {
	channel := Connect()
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(q)

	err = channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Testing..."),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Published Message to Queue")
}
