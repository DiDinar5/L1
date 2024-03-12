package main

import "fmt"

// NotificationSender интерфейс для отправки уведомлений
type NotificationSender interface {
	SendNotification(message string)
}

// EmailNotificationSender реализация отправки уведомлений по email
type EmailNotificationSender struct{}

func (s *EmailNotificationSender) SendNotification(message string) {
	fmt.Println("Sending email:", message)
}

// ExternalService реализация стороннего сервиса для отправки уведомлений
type ExternalService struct{}

func (s *ExternalService) SendMessageToUser(message, userContact string) {
	fmt.Printf("Using external service to send '%s' to '%s'\n", message, userContact)
}

// ExternalServiceAdapter адаптер для ExternalService, чтобы использовать его как NotificationSender
type ExternalServiceAdapter struct {
	Service     *ExternalService
	UserContact string
}

func (a *ExternalServiceAdapter) SendNotification(message string) {
	a.Service.SendMessageToUser(message, a.UserContact)
}

func main() {
	emailSender := &EmailNotificationSender{}
	emailSender.SendNotification("Hello via email!")

	// Создаем экземпляр стороннего сервиса
	externalService := &ExternalService{}
	// Адаптируем сторонний сервис к интерфейсу NotificationSender
	adapter := &ExternalServiceAdapter{
		Service:     externalService,
		UserContact: "user@example.com",
	}
	adapter.SendNotification("Hello via external service!")
}
