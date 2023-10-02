package mq

import (
	"fmt"

	"github.com/bottlehub/unboard/boards/configs"
	"github.com/bottlehub/unboard/boards/internal"
	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	connection, err := amqp.Dial(configs.EnvRabbitMQ())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to RabbitMQ Instance Successfully")

	channel, err := connection.Channel()
	internal.Handle(err)

	return channel
}
