package factory

import "fmt"

const (
	EmailMessage int = 1
	SmsMessage   int = 2
)

type Notification struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type Notifier interface {
	CreateNotification(Notification) string
}

type SmsNotification struct {
}

func (s *SmsNotification) CreateNotification(notification Notification) string {
	return fmt.Sprintf("Sms body = { %s }   Sent to PhoneNo { %s }", notification.Message, notification.To)
}

type EmailNotification struct{}

func (e *EmailNotification) CreateNotification(notification Notification) string {
	return fmt.Sprintf("Email Body = { %s }  Sent to Email { %s } ", notification.Message, notification.To)
}

func Factory(messageType int) (Notifier, error) {

	switch messageType {
	case EmailMessage:
		{
			return &EmailNotification{}, nil
		}
	case SmsMessage:
		{
			return &SmsNotification{}, nil
		}
	default:
		{
			return nil, fmt.Errorf("Unsupported message type")
		}

	}
}
