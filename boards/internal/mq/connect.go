package mq

import (
	"fmt"

	"github.com/bottlehub/unboard/boards/configs"
	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	connection, err := amqp.Dial(configs.EnvRabbitMQ())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to RabbitMQ Instance Successfully")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}
