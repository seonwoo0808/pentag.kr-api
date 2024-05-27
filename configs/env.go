package configs

import "os"

type EnvType struct {
	CaptchaSecret string
	SMTPUser string
	SMTPHost string
	SMTPPort string
	SMTPPassword string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

var Env EnvType

func loadEnv(key string) string{
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	panic("Environment variable " + key + " is not set")
}
	
func LoadAllEnv() {
	Env.CaptchaSecret = loadEnv("CAPTCHA_SECRET")
	Env.SMTPUser = loadEnv("SMTP_USER")
	Env.SMTPHost = loadEnv("SMTP_HOST")
	Env.SMTPPort = loadEnv("SMTP_PORT")
	Env.SMTPPassword = loadEnv("SMTP_PASSWORD")
	Env.DBHost = loadEnv("DB_HOST")
	Env.DBPort = loadEnv("DB_PORT")
	Env.DBUser = loadEnv("DB_USER")
	Env.DBPassword = loadEnv("DB_PASSWORD")
	Env.DBName = loadEnv("DB_NAME")
}