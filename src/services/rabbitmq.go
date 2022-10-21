package services

import (
	"context"
	"log"
	"os"

	jsonEncoder "encoding/json"

	"github.com/Fonzeca/Chatline/src/db/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupRabbitMq() (*amqp.Channel, func()) {
	// Create a new RabbitMQ connection.

	connectRabbitMQ, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		panic(err)
	}

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		connectRabbitMQ.Close()
		panic(err)
	}

	GlobalChannel = channelRabbitMQ
	return channelRabbitMQ, func() { connectRabbitMQ.Close(); channelRabbitMQ.Close() }
}

func ProcessData(commentMq model.ComentarioView) error {
	commentBytes, _ := jsonEncoder.Marshal(commentMq)
	err := GlobalChannel.PublishWithContext(context.Background(), "carmind", "notification.comment.chatline.preparing", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        commentBytes,
	})
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
