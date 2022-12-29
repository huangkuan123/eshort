package bootstrap

import "eshort/pkg/rabbitmq"

func SetUpRabbitMQ() {
	rabbitmq.ConnectRabbitMQ()
}
