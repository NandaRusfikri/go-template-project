package dto

type SendEmail struct {
	To         []string
	Cc         []string
	Bcc        []string
	Subject    string
	BodyType   string
	Body       string
	Attachment []string
}
type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
	Name     string
}
