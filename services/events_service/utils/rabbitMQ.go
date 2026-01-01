package utils

import (
	"fmt"
	"os"
)

func BuildRabbitURL() string {
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	if port == "" {
		port = "5672"
	}

	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		user,
		pass,
		host,
		port,
	)
}
