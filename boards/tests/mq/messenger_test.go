package mq_test

import (
	"log"
	"os"
	"testing"
)

func TestPublish(t *testing.T) {
	env, err := os.LookupEnv("RABBITMQ")
	if !err {
		log.Fatal("Error loading .env file: ", err)
	}
	log.Println(env)
}
