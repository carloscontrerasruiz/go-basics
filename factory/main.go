package main

import "fmt"

//SMS Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

//SMS
type SMSNotification struct {
}

type SMSNotificationSender struct {
}

//SMS SENDER
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "TWILIO"
}

//SMS FACTORY
func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification using SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

//EMAIL
type EmailNotification struct {
}

type EmailNotificationSender struct {
}

//EMAIL SENDER
func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "GMAIL"
}

//EMAIL FACTORY
func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification using SMTP")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

//============= FACTORY==================

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "EMAIL":
		return &EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("Type invalid")
	}
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	getMethod(smsFactory)

	sendNotification(emailFactory)
	getMethod(emailFactory)
}
