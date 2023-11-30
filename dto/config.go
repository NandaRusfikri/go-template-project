package dto

type ConfigEnvironment struct {
	DbUser    string
	DbPass    string
	DbHost    string
	DbPort    string
	DbName    string
	DbSslmode string

	SmtpHost     string
	SmtpPort     string
	SmtpEmail    string
	SmtpPassword string
	SmtpName     string

	Timezone    string
	Version     string
	RestPort    string
	GoEnv       string
	SwaggerHost string
	JwtSecret   string
}
