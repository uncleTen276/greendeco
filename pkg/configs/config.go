package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host string `envconfig:"APP_HOST" default:"localhost"`
		PORT string `envconfig:"APP_PORT" default:":8080"`
	}
	Database struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Name     string `envconfig:"DB_NAME" default:"greendeco"`
		Port     string `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" default:"postgres"`
		Password string `envconfig:"DB_PASSWORD" default:"postgres"`
		SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
	}
	Auth struct {
		JWTSecret        string `envconfig:"JWT_SECRET" default:"token-secret"`
		TokenExpire      int    `envconfig:"TOKEN_EXPIRE" default:"60"`
		ShortTokenExpire int    `envconfig:"SHORT_TOKEN_EXPIRE" default:"15"`
	}
	SMTP struct {
		Email             string `envconfig:"SMTP_EMAIL" default:""`
		Password          string `envconfig:"SMTP_PASSWORD" default:""`
		LinkResetPassword string `envconfig:"SMTP_LINK_RESET_PSW" default:"localhost:8080"`
	}
	Storage struct {
		Firebase struct {
			Type                    string `envconfig:"FIREBASE_TYPE" default:"" json:"type"`
			ProjectID               string `envconfig:"FIREBASE_PROJECT_ID" default:"" json:"project_id"`
			PrivateKeyID            string `envconfig:"FIREBASE_PRIVATE_KEY_ID" default:"" json:"private_key_id"`
			PrivateKey              string `envconfig:"FIREBASE_PRIVATE_KEY" default:"" json:"private_key"`
			ClientEmail             string `envconfig:"FIREBASE_CLIENT_EMAIL" default:"" json:"client_email"`
			ClientID                string `envconfig:"FIREBASE_CLIENT_ID" default:"102751571694782723192" json:"client_id"`
			AuthURI                 string `envconfig:"FIREBASE_AUTH_URI" default:"https://accounts.google.com/o/oauth2/auth" json:"auth_uri"`
			TokenURI                string `envconfig:"FIREBASE_TOKEN_URI" default:"https://oauth2.googleapis.com/token" json:"token_uri"`
			AuthProviderX509CertURL string `envconfig:"FIREBASE_AUTH_PROVIDER_X509_CERT_URL" default:"https://www.googleapis.com/oauth2/v1/certs" json:"auth_provider_x509_cert_url"`
			ClientX509CertURL       string `envconfig:"CLIENT_X509_CERT_URL" default:"" json:"client_x509_cert_url"`
			UniverseDomain          string `envconfig:"UNIVERSE_DOMAIN" default:"googleapis.com" json:"universe_domain"`
		}
		Bucket string `envconfig:"FIREBASE_BUCKET" default:""`
	}
	VnPay struct {
		Secret     string `envconfig:"VNPAY_SECRET" default:"XNBCJFAKAZQSGTARRLGCHVZWCIOIGSHN"`
		TmnCode    string `envcofig:"VNPAY_TMNCODE" default:"CGXZLS0Z"`
		ReturnUrl  string `envconfig:"VNPAY_RETURN_URL" default:"http://localhost:8080/api/v1/payment/vnpay_return"`
		CancelUrl  string `envconfig:"VNPAY_CANCEL_URL" default:"http://localhost:3000/user/order/"`
		SuccessUrl string `envconfig:"VNPAY_SUCCESS_URL" default:"http://localhost:3000/payment-result/success"`
		ErrorUrl   string `envconfig:"VNPAY_ERROR_URL" default:"http://localhost:3000/payment-result/error"`
	}
	PayPal struct {
		ClientId   string `envconfig:"PAYPAL_CLIENT" default:""`
		SecretKey  string `envconfig:"PAYPAL_SECRET_KEY" default:""`
		ReturnUrl  string `envconfig:"PAYPAL_RETURN_URL" default:"http://localhost:3000/payment/error" `
		SuccessUrl string `envconfig:"PAYPAL_SUCCESS_URL" default:"http://localhost:3000/payment-result/success"`
	}

	ExchangeMoneyApi struct {
		Url string `envconfig:"EXCHANGE_MONEY_API" default:""`
	}
}

var appConfig = &Config{}

func AppConfig() *Config {
	return appConfig
}

func LoadConfig() error {
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	return nil
}
