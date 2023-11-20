package pkg

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-template-project/dto"
	"os"
)

func LoadEnvironment(path string) (config dto.ConfigEnvironment) {
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read environment variables from docker-compose.yml
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_SSLMODE = os.Getenv("DB_SSLMODE")

	config.SMTP_HOST = os.Getenv("SMTP_HOST")
	config.SMTP_PORT = os.Getenv("SMTP_PORT")
	config.SMTP_NAME = os.Getenv("SMTP_NAME")
	config.SMTP_EMAIL = os.Getenv("SMTP_EMAIL")
	config.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

	config.TIMEZONE = os.Getenv("TIMEZONE")
	config.VERSION = os.Getenv("VERSION")
	config.REST_PORT = os.Getenv("REST_PORT")
	config.GO_ENV = os.Getenv("GO_ENV")
	config.SWAGGER_HOST = os.Getenv("SWAGGER_HOST")
	config.JWT_SECRET = os.Getenv("JWT_SECRET")

	return config
}
