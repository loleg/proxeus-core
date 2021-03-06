package service

import (
	"github.com/ProxeusApp/proxeus-core/sys/email"
)

type (
	// EmailService is an interface that provides functions to send emails
	EmailService interface {

		// Send dispatches an email. The body can contain html tags.
		// The sender of email will be the default EmailFrom value that is configured in the settings.
		Send(emailTo, subject, body string) error

		// SendFrom dispatches an email. The body can contain html tags.
		// Additionally specify the sender of the email with the emailFrom parameter
		SendFrom(emailFrom, emailTo, subject, body string) error
	}
	DefaultEmailService struct {
	}
)

func NewEmailService() EmailService {
	return &DefaultEmailService{}
}

// Send dispatches an email. The body can contain html tags.
// The sender of email will be the default EmailFrom value that is configured in the settings.
func (me *DefaultEmailService) Send(emailTo, subject, body string) error {
	settings, err := settingsDB().Get()
	if err != nil {
		return err
	}

	return me.SendFrom(settings.EmailFrom, emailTo, subject, body)
}

// SendFrom dispatches an email. The body can contain html tags.
// Additionally specify the sender of the email with the emailFrom parameter
func (me *DefaultEmailService) SendFrom(emailFrom, emailTo, subject, body string) error {
	mail := &email.Email{From: emailFrom, To: []string{emailTo}, Subject: subject, Body: body}

	return emailSender().Send(mail)
}
