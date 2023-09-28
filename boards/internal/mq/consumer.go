package mq

import (
	"fmt"
)

func Consume(queue string) []byte {
	channel := Connect()
	defer channel.Close()

	msgs, err := channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	concur := make(chan []byte)
	go func() {
		for d := range msgs {
			concur <- d.Body
		}
	}()

	fmt.Println(" [*] - Waiting for Messages")
	return <-concur
}
