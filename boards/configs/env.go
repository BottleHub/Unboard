package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Helps retrieve the MongoDB URI fron the env file
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Fatal("Error loading .env file: ", err)
		}
	}
	return os.Getenv("MONGOURI")
}

func EnvRabbitMQ() string {
	env, err := os.LookupEnv("RABBITMQ")
	if !err {
		err := godotenv.Load()
		if err != nil {
			err := godotenv.Load(".env.test")
			if err != nil {
				log.Fatal("Error loading .env file: ", err)
			}
		}
		return os.Getenv("RABBITMQ")
	}
	log.Println(env)
	return env
}
