package mq

import (
	"fmt"
)

func Consume(queue string) {
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

	delay := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received: %s\n", d.Body)
		}
	}()

	fmt.Println(" [*] - Waiting for Messages")
	<-delay
}
