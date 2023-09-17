package mq

import (
	"fmt"
	"log"
	"os"

	"github.com/bottlehub/unboard/boards/configs"
	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	url := configs.EnvRabbitMQ()
	if url == "" {
		env, err := os.LookupEnv("RABBITMQ")
		if !err {
			log.Fatal("Error loading .env file: ", err)
		}
		url = env
	}

	connection, err := amqp.Dial(url)
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
