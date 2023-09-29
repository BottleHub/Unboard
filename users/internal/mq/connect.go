package mq

import (
	"fmt"

	"github.com/bottlehub/unboard/users/configs"
	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	connection, err := amqp.Dial(configs.EnvRabbitMQ())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to RabbitMQ Insatnce Successfully")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}
