package dto

type ConfigEnvironment struct {
	Database ConfigDatabase
	SMTP     ConfigSMTP
	App      ConfigApp
}
type ConfigApp struct {
	Timezone    string
	Version     string
	RestPort    string
	GoENV       string
	SwaggerHost string
	JwtSecret   string
}
type ConfigDatabase struct {
	User    string
	Pass    string
	Host    string
	Port    string
	Name    string
	SSLmode string
}
type ConfigSMTP struct {
	Host     string
	Port     string
	Email    string
	Password string
	Name     string
}
