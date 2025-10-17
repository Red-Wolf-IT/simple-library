package main

import "fmt"

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю на email %s сообщение: %s\n", e, message)
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Отправляю SMS на номер %s сообщение: %s\n", s, message)
}

type Notifier interface {
	Notify(message string)
}
