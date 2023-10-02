package mq

import (
	"fmt"

	"github.com/bottlehub/unboard/boards/internal"
)

func Consume(queue string, fn func(string)) {
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
	internal.Handle(err)

	delay := make(chan bool)
	go func() {
		for d := range msgs {
			s := fmt.Sprintf("%s\n", d.Body)
			fn(s)
		}
	}()

	fmt.Println(" [*] - Waiting for Messages")
	<-delay
}
