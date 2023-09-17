package mq

import "fmt"

func Consume() {
	channel := Connect()
	defer channel.Close()

	msgs, err := channel.Consume(
		"TestQueue",
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

	concur := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received: %s\n", d.Body)
		}
	}()

	fmt.Println(" [*] - Waiting for Messages")
	<-concur
}
