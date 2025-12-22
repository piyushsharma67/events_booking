package service

import "fmt"

type MockNotifier struct {
	SentMessages []Notification
}

type Notification struct {
	To      string
	Subject string
	Body    string
}

func (m *MockNotifier) SendNotification(to string, subject string, body string) error {
	m.SentMessages = append(m.SentMessages, Notification{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	fmt.Printf("Mock send to %s with subject %s\n", to, subject) // optional debug
	return nil
}