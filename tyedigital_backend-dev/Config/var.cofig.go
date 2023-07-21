package config

import (
	"database/sql"
	"os"

	model "github.com/Kushmanda-Tech/tyedigital_backend/Model"
)

type DBConfig struct {
	DB_PORT                string
	DB_USERNAME            string
	DB_PASSWORD            string
	DB_HOST                string
	SERVER_SUB_DOMAIN_NAME string
	SERVER_DOMAIN_NAME     string
	SERVER_NAME            string
}

func DbConfig() DBConfig {
	c := DBConfig{
		DB_PORT:                os.Getenv("DB_PORT"),
		DB_USERNAME:            os.Getenv("DB_USERNAME"),
		DB_PASSWORD:            os.Getenv("DB_PASSWORD"),
		DB_HOST:                os.Getenv("DB_HOST"),
		SERVER_SUB_DOMAIN_NAME: os.Getenv("SERVER_SUB_DOMAIN_NAME"),
		SERVER_DOMAIN_NAME:     os.Getenv("SERVER_DOMAIN_NAME"),
		SERVER_NAME:            os.Getenv("SERVER_SUB_DOMAIN_NAME") + os.Getenv("SERVER_DOMAIN_NAME"),
	}
	return c
}

var (
	DB                     *sql.DB
	SERVER_SUB_DOMAIN_NAME string
	SERVER_DOMAIN_NAME     string
	DB_NAME                = "TYEDIGITAL"
	SERVER_PORT            = "6060"
	INGRESS_PATH           = "api"
	VERIFY_EMAIL_API       = "https://" + DbConfig().SERVER_NAME + "/" + INGRESS_PATH + "/verify_email?email="
	SERVER_CERT_FILE       = "/etc/letsencrypt/live/" + DbConfig().SERVER_NAME + "/fullchain.pem"
	SERVER_KEY_FILE        = "/etc/letsencrypt/live/" + DbConfig().SERVER_NAME + "/privkey.pem"

	NO_REPLY_EMAIL_PASSWORS = "Luv974@Ku$h"
	EMAIL_SEND_FROM         = "no.reply@kushmanda.tech"
	SMTP_SERVER_ADDRESS     = "smtp.office365.com"
	SMTP_PORT               = "587"
)

const (
	FIREBASE_AUTH_KEY = "key=AAAARciEovc:APA91bEAFNfEDPcCnsazxkz7ja-vYr4wbwaJtjD0Vk4lZDvTc8K06JKfAAteSjNbaymYPc93AKxIF70Wm32Md6lgQGnTbfiX4erX62ADYI93LnA8uWpOcyGiuTcUPkg3vBRT_NekKprm"
	API_VERSION       = "/V2.0"
	SecretKeyHex      = "b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2"
	PublicKeyHex      = "1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2"
	Audience          = "kushmanda.tech"
	Identifier        = "kushmanda.tech"
	Issuer            = "kushmanda.tech"
	Subject           = "REST_API"
)

var Input = model.InputStruct{}
