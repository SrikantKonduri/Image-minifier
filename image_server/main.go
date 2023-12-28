package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"image_server/databases"
	"image_server/utils"

	"github.com/joho/godotenv"

	amqp "github.com/rabbitmq/amqp091-go"
)

func initEnv() {
	er := godotenv.Load()
	if er != nil {
		utils.FailOnError(er, "Failed to read from .env")
	}
}

func main() {
	db := databases.ConnectDB()
	db = db
	initEnv()
	msgQueue := fmt.Sprintf("amqp://guest:guest@%s:%s/", os.Getenv("MSG_QUEUE_HOST"), os.Getenv("MSG_QUEUE_PORT"))
	conn, err := amqp.Dial(msgQueue)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("MSG_QUEUE_NAME"), // name
		false,                       // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		nil,                         // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	q = q
	msgs, err := ch.Consume(
		os.Getenv("MSG_QUEUE_NAME"), // queue
		"",                          // consumer
		true,                        // auto-ack
		false,                       // exclusive
		false,                       // no-local
		false,                       // no-wait
		nil,                         // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			_, err := strconv.Atoi(string(d.Body))
			if err != nil {
				log.Printf("Error converting message to int: %v", err)
			}
			// handlers.HandleProduct(db, int64(value))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
