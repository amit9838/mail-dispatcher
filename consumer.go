package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	smtpHost := "localhost"
	smtpPort := "1025"

	for recipient := range ch {
		fmt.Printf("Emailing %d %s at %s\n", id, recipient.Name, recipient.Email)

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\nHello %s,\r\nThis is a test email\r\n", recipient.Email, recipient.Name)
		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient)
		fmt.Printf("Worker %d: Sending email to %s \n", id, recipient.Email)

		if err != nil {
			fmt.Printf("Worker %d: Failed to execute template: %v\n", id, err)
			continue
		}

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "amit@gmail.com", []string{recipient.Email}, []byte(msg))

		if err != nil {
			fmt.Printf("Worker %d: Failed to send email to %s: %v\n", id, recipient.Email, err)
		}

		// Simulate email sending delay
		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Email sent to %s\n", id, recipient.Email)
	}
}
