package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Publish(queue string, message string) (int, error) {
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
	if err != nil {
		ch := make(chan string)
		go panic(err)
		<-ch
		return 1, err
	}

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
	if err != nil {
		ch := make(chan string)
		go panic(err)
		<-ch
		return 2, err
	}

	fmt.Println("Published Message to Queue")
	return 0, err
}
