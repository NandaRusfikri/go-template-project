package pkg

import (
	log "github.com/sirupsen/logrus"
	"go-template-project/dto"
	"gopkg.in/gomail.v2"
	"strconv"
)

type SMTP struct {
	Config *dto.ConfigSMTP
}

func InitEmail(config *dto.ConfigSMTP) *SMTP {
	return &SMTP{
		Config: config,
	}
}

func (e *SMTP) SendEmail(params dto.SendEmail) error {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.Config.Email)
	mailer.SetHeader("To", params.To...)

	if params.Cc != nil && len(params.Cc) > 0 {
		mailer.SetHeader("Cc", params.Cc...)
	}

	if params.Bcc != nil && len(params.Bcc) > 0 {
		mailer.SetHeader("Bcc", params.Bcc...)
	}

	mailer.SetHeader("Subject", params.Subject)
	if params.BodyType == "" {
		params.BodyType = "text/html"
	}
	mailer.SetBody(params.BodyType, params.Body)

	if params.Attachment != nil && len(params.Attachment) > 0 {
		for _, path := range params.Attachment {
			mailer.Attach(path)
		}
	}

	Port, err := strconv.Atoi(e.Config.Port)
	if err != nil {
		log.Errorln("SMTP Port is not valid ", err.Error())
	}

	dialer := gomail.NewDialer(
		e.Config.Host,
		Port,
		e.Config.Email,
		e.Config.Password,
	)
	dialer.TLSConfig.InsecureSkipVerify = true

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Errorln("❌ Email - Send - error: ", err)
		return err
	}

	return nil
}
