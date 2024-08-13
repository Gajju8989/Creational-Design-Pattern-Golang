package main

import (
	"design_patterns/factory"
	"fmt"
)

func main() {
	// Email Payload
	emailNotificationPayload := factory.Notification{
		To:      "example@test.com",
		Message: "Hello",
	}

	method, err := factory.Factory(factory.EmailMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := method.CreateNotification(emailNotificationPayload)
	fmt.Println(resp)

	// SMS Payload
	smsNotificationPayload := factory.Notification{
		To:      "7999999999",
		Message: "Hello",
	}

	method, err = factory.Factory(factory.SmsMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp = method.CreateNotification(smsNotificationPayload)
	fmt.Println(resp)
}
