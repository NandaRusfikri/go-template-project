package dto

type ConfigEnvironment struct {
	DB_USER    string
	DB_PASS    string
	DB_HOST    string
	DB_PORT    string
	DB_NAME    string
	DB_SSLMODE string

	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_EMAIL    string
	SMTP_PASSWORD string
	SMTP_NAME     string

	TIMEZONE     string
	VERSION      string
	REST_PORT    string
	GO_ENV       string
	SWAGGER_HOST string
	JWT_SECRET   string
}
