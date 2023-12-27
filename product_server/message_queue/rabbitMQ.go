package messagequeue

import (
	"context"
	"fmt"
	"log"
	"os"
	"product_server/utils"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToQueue() (*amqp.Channel, context.Context, error) {
	queueHost := os.Getenv("MSG_QUEUE_HOST")
	queuePort := os.Getenv("MSG_QUEUE_PORT")
	queueName := os.Getenv("MSG_QUEUE_NAME")
	msgQueue := fmt.Sprintf("amqp://guest:guest@%s:%s/", queueHost, queuePort)
	conn, err := amqp.Dial(msgQueue)
	// status := utils.FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	// status = utils.FailOnError(err, "Failed to open a channel")
	// defer ch.Close()
	if err != nil {
		return nil, nil, err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	q = q
	if err != nil {
		return nil, nil, err
	}
	// utils.FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	cancel = cancel

	return ch, ctx, nil

}

func ProduceMessage(productId int64, ch *amqp.Channel, ctx context.Context, qName string) error {
	// body := 11
	bodyBytes := []byte(strconv.FormatInt(productId, 10))
	err := ch.PublishWithContext(ctx,
		"",    // exchange
		qName, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bodyBytes,
		})
	if err != nil {
		utils.FailOnError(err, "Failed to publish a message")
		return err
	}
	log.Printf(" [x] Sent %s\n", productId)
	return nil
}
