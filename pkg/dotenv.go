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
	config.DbHost = os.Getenv("DB_HOST")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbUser = os.Getenv("DB_USER")
	config.DbName = os.Getenv("DB_NAME")
	config.DbPass = os.Getenv("DB_PASS")
	config.DbSslmode = os.Getenv("DB_SSLMODE")

	config.SmtpHost = os.Getenv("SMTP_HOST")
	config.SmtpPort = os.Getenv("SMTP_PORT")
	config.SmtpName = os.Getenv("SMTP_NAME")
	config.SmtpEmail = os.Getenv("SMTP_EMAIL")
	config.SmtpPassword = os.Getenv("SMTP_PASSWORD")

	config.Timezone = os.Getenv("TIMEZONE")
	config.Version = os.Getenv("VERSION")
	config.RestPort = os.Getenv("REST_PORT")
	config.GoEnv = os.Getenv("GO_ENV")
	config.SwaggerHost = os.Getenv("SWAGGER_HOST")
	config.JwtSecret = os.Getenv("JWT_SECRET")

	return config
}
