package pkg

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
	Name     string
}

type SMTP struct {
	Config *SMTPConfig
}

func InitEmail(config *SMTPConfig) *SMTP {
	return &SMTP{
		Config: config,
	}
}

func (e *SMTP) SendEmail(to []string, cc []string, bcc []string, subject string, bodyType string, body string, attachment []string) error {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.Config.Email)
	mailer.SetHeader("To", to...)

	if cc != nil && len(cc) > 0 {
		mailer.SetHeader("Cc", cc...)
	}

	if bcc != nil && len(bcc) > 0 {
		mailer.SetHeader("Bcc", bcc...)
	}

	mailer.SetHeader("Subject", subject)
	if bodyType == "" {
		bodyType = "text/html"
	}
	mailer.SetBody(bodyType, body)

	if attachment != nil && len(attachment) > 0 {
		for _, path := range attachment {
			mailer.Attach(path)
		}
	}

	dialer := gomail.NewDialer(
		e.Config.Host,
		e.Config.Port,
		e.Config.Email,
		e.Config.Password,
	)
	//dialer.TLSConfig.InsecureSkipVerify = true

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Errorln("❌ Email - Send - error: ", err)
		return err
	}

	return nil
}
