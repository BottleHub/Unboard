package mq

import (
	"fmt"
	"log"

	"github.com/bottlehub/unboard/boards/configs"
	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	connection, err := amqp.Dial(configs.EnvRabbitMQ())
	if err != nil {
		log.Fatal(err, configs.EnvRabbitMQ())
	}

	fmt.Println("Connected to RabbitMQ Insatnce Successfully")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}
