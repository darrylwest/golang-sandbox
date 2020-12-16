package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"time"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "darryl.west@raincitysoftware.com")
	m.SetHeader("To", "7752508168@messaging.sprintpcs.com"
	// m.SetHeader("Subject", "Test go lang gomail email")
	msg := fmt.Sprintf("This is a test message sent at %v", time.Now().Unix())
	m.SetBody("text/plain", msg)

	d := gomail.NewPlainDialer("smtp.example.com", 587, "user", "123456")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
