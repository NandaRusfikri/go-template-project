package pkg

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go-template-project/dto"
	"os"
)

func LoadConfig(path string) (config dto.ConfigEnvironment) {
	err := godotenv.Load(path)
	if err != nil {
		log.Errorln("Error loading .env file ", err)
	}
	config.Database = dto.ConfigDatabase{
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		User:    os.Getenv("DB_USER"),
		Name:    os.Getenv("DB_NAME"),
		Pass:    os.Getenv("DB_PASS"),
		SSLmode: os.Getenv("DB_SSLMODE"),
	}
	config.SMTP = dto.ConfigSMTP{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Name:     os.Getenv("SMTP_NAME"),
		Email:    os.Getenv("SMTP_EMAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}
	config.App = dto.ConfigApp{
		Timezone:    os.Getenv("TIMEZONE"),
		Version:     os.Getenv("VERSION"),
		RestPort:    os.Getenv("REST_PORT"),
		GoENV:       os.Getenv("GO_ENV"),
		SwaggerHost: os.Getenv("SWAGGER_HOST"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}

	return config
}
