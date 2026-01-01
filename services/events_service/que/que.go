package que

import (
	"context"

	"github.com/streadway/amqp"
)

type QuePublisher struct {
	ch    *amqp.Channel
	queue string
}

func NewEventsPublisher(conn *amqp.Connection, queue string) (*QuePublisher, error) {
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

	return &QuePublisher{
		ch:    ch,
		queue: queue,
	}, nil
}

func (p *QuePublisher) Publish(ctx context.Context, message []byte) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return p.ch.Publish(
			"",
			p.queue,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        message,
			},
		)
	}
}
