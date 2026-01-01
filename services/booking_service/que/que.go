package que

import (
	"context"
	"log"

	"github.com/piyushsharma67/events_booking/services/booking_service/repository"
	"github.com/streadway/amqp"
)

type QueConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
	handler func(amqp.Delivery) error
	repository repository.SeatsRepository
}

func NewBookingConsumer(conn *amqp.Connection, queue string,handler func(amqp.Delivery) error) (*QueConsumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &QueConsumer{
		conn:    conn,
		channel: ch,
		queue:   queue,
		handler: handler,
	}, nil
}

func (c *QueConsumer) Start(ctx context.Context) error {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("booking consumer shutting down")
			return nil

		case msg := <-msgs:
			if err := c.handler(msg); err != nil {
				log.Println("handler error:", err)
				msg.Nack(false, true) // retry
				continue
			}

			msg.Ack(false)
		}
	}
}
