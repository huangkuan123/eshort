package rabbitmq

import (
	"eshort/pkg/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var RabbitMQConn *amqp.Connection

// 后期需改为用连接池，Connection进池，Channel进池
func ConnectRabbitMQ() {
	var err error
	//我的虚拟机地址
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		config.GetString("rabbitmq.user"),
		config.GetString("rabbitmq.password"),
		config.GetString("rabbitmq.host"),
		config.GetString("rabbitmq.port"),
		config.GetString("rabbitmq.vhost"),
	)
	RabbitMQConn, err = amqp.Dial(url)
	FailOnError(err, "RabbitMQ 连接失败")
}

func GetChannel() *amqp.Channel {
	channel, err := RabbitMQConn.Channel()
	FailOnError(err, "打开频道失败")
	return channel
}

func GetQueueDeclare(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare( //声明一个队列(没有时会自动创建)
		queueName, // name 队列名称
		false,     // durable 是否持久化到磁盘中。如果为true，重启后，该队列是否也会恢复。autoDelete 参数受此影响，恢复时，也会带上这个参数。
		false,     // autoDelete delete when unused 是否自动删除：当未使用时删除
		false,     // exclusive 是否独占，true则表示这个队列只能供一个消费者使用
		false,     // noWait 为true时，将默认该队列已经声明了。如果满足现有队列的条件或试图从其他连接修改现有队列，则会出现通道异常
		nil,       // args 其他参数
	)
	FailOnError(err, "队列创建失败")
	return q
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
