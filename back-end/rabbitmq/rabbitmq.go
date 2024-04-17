package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	user := os.Getenv("RABBITMQ_DEFAULT_USER")
	pass := os.Getenv("RABBITMQ_DEFAULT_PASS")
	port := os.Getenv("RABBITMQ_PORT")
	url := fmt.Sprintf("amqp://%s:%s@rabbitmq:%s/", user, pass, port)

	var conn *amqp.Connection

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		var err error              
		conn, err = amqp.Dial(url) 
		if err != nil {
			log.Println("Failed to connect to RabbitMQ, retrying...")
			time.Sleep(time.Second * 5)
			continue
		}

		channel, err := conn.Channel()
		if err != nil {
			log.Println("Failed to open a channel, retrying...")
			conn.Close() 
			time.Sleep(time.Second * 5)
			continue
		}

		return &RabbitMQ{
			conn:    conn,
			channel: channel,
		}, nil
	}

	// Close the connection if maximum retries reached
	if conn != nil {
		conn.Close()
	}

	return nil, fmt.Errorf("failed to connect to RabbitMQ after %d retries", maxRetries)
}

func (rmq *RabbitMQ) Close() {
	if rmq.channel != nil {
		rmq.channel.Close()
	}
	if rmq.conn != nil {
		rmq.conn.Close()
	}
}

func (rmq *RabbitMQ) DeclareQueue(queueName string) error {
	_, err := rmq.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return err
}

func (rmq *RabbitMQ) Publish(queueName string, body []byte) error {
	return rmq.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func (rmq *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	return rmq.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}
