package mq

import (
	"Phase3/week2/day3/go-pooling-message-broker/internal/app/interfaces"
	"fmt"

	"github.com/streadway/amqp"
)

type MessageQueue struct {
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

var _ interfaces.MessageQueue = &MessageQueue{}

// Init rabbit MQ
func NewMessageQueue(url string) (*MessageQueue, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	return &MessageQueue{Connection: conn, Channel: ch}, nil
}

func (mq *MessageQueue) Close() {
	mq.Channel.Close()
	mq.Connection.Close()
}

func (mq *MessageQueue) PublishItemCreated(message []byte) error {
	return mq.publish("item.created", message)
}

func (mq *MessageQueue) PublishItemUpdated(message []byte) error {
	return mq.publish("item.updated", message)
}

func (mq *MessageQueue) PublishItemDeleted(id int) error {
	msg := fmt.Sprintf("Item deleted: %d", id)
	return mq.publish("item.deleted", []byte(msg))
}

func (mq *MessageQueue) publish(routingKey string, message []byte) error {
	return mq.Channel.Publish(
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}
