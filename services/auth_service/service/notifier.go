package service

import (
	"encoding/json"
	"errors"

	"github.com/streadway/amqp"
)

type Notifier interface {
	SendNotification(to string, subject string, body string) error
}

type MessageBrokerService struct {
	ch       *amqp.Channel
	que      string
	confirms <-chan amqp.Confirmation
}

func NewRabbitMQNotifier(conn *amqp.Connection, queue string) (*MessageBrokerService, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Enable publisher confirms
	if err := ch.Confirm(false); err != nil {
		return nil, err
	}

	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

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

	return &MessageBrokerService{ch: ch, que: queue, confirms: confirms}, nil
}

type EmailNotification struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (r *MessageBrokerService) SendNotification(to, subject, body string) error {
	msg := EmailNotification{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = r.ch.Publish(
		"",
		r.que,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         data,
		},
	)
	if err != nil {
		return err
	}

	// 🔴 THIS IS THE ACTUAL GUARANTEE
	confirm := <-r.confirms
	if !confirm.Ack {
		return errors.New("rabbitmq did not acknowledge message")
	}

	return nil
}
