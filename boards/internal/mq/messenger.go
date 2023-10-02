package mq

import (
	"fmt"

	"github.com/bottlehub/unboard/boards/internal"
	"github.com/streadway/amqp"
)

func Publish(queue string, message string) error {
	channel := Connect()
	defer channel.Close()

	q, err := channel.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	internal.Handle(err)

	fmt.Println(q)

	err = channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(message),
		},
	)
	internal.Handle(err)

	fmt.Println("Published Message to Queue")
	return err
}
