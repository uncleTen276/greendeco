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
		Email             string `envconfig:"SMTP_EMAIL" default:"kristiannguyen276@gmail.com"`
		Password          string `envconfig:"SMTP_PASSWORD" default:"cdrplghspkujcvkf"`
		LinkResetPassword string `envconfig:"SMTP_LINK_RESET_PSW" default:"localhost:8080"`
	}
	Storage struct {
		Firebase struct {
			Type                    string `envconfig:"FIREBASE_TYPE" default:"service_account" json:"type"`
			ProjectID               string `envconfig:"FIREBASE_PROJECT_ID" default:"greendeco-2726b" json:"project_id"`
			PrivateKeyID            string `envconfig:"FIREBASE_PRIVATE_KEY_ID" default:"0550d241c94c942f76786c3257c018a34ca86a04" json:"private_key_id"`
			PrivateKey              string `envconfig:"FIREBASE_PRIVATE_KEY" default:"-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC9NE5wM5SI4vrX\nWXKUnZ/rFOa+a5JuRysz14sTIsD5qUTbYx/fJjR/tgI1PdXBHvUkGkXErPDrG7h0\nIzmE6IVS/Ruade0AR6hpjoyC++zmvpXx/evjKNcp9u7SjZ9w6+8WzP0FtarAG3Ku\n5syHSCA/jLJIXjjEKKHM+Pfp1O7qjPxeII5ceYqHY6zvApDKlCBilraKAL654Eck\nYe7I5S3TcjZqQ8Lxag19dPupfKOEXO0SoEHoH9QB5+YJ6qtwpEF4hGszwKWWMdFy\n9DEmGniJlu0ojcltjNXGn3D5SGJfCeTyvX5Uf2wqam9I92F6RGSYZCYV0IKzw6x5\nX1lDXXl5AgMBAAECggEAGxbMmN2XCXxErCWHNVurZFCN8uXa8Bq/cZY4OB6p4MpJ\nQgdGA8CSvqrpU68ilhpfwHIEhYgPRvzP4sznn4MHvQiLywh8TNB1qTClfIbOSyvZ\nXcLcYvMwzzYmyRL0KN+Wz+hnMknphG90SDIM8YeJmnUDUj+EmCKLJpo01eVMlFx+\nackduFE2iF9RosR+WUUY1xfQMAUDeNj5AmKwMcOEljKSohHShG9oG4BgHiDXjU3j\nRrB4eLt+2HXuy5LEbgrQjaDyQ2E4VshDkpa3EqDgRgL3GHPmk+LDjL8E/WA7cc16\n0ziKCB85ubhgdxwmfXFcVpTj9ciSs0fbVfpwFaL4IwKBgQDoKqEbhYNTrG2UrelN\nIuGQfQRit1z48887klQSD/+6VrfgtMz9JUg93R8KCJ/ONmgXNbc1xg0xIw4uLeWF\ndsVh/nmSOOq+dM4Ju0+T4D1DenMQ/2jDtimKG3OhsOqFepOC4FH0lnzguJkx5Ufh\nlvqeWKLZsJ9MAtJnQC0xA4dgmwKBgQDQoJ/FA1gCZtAd4Z011pNXyMeGpp9UiQbz\nK3qLxnM560FxwHSkZ7qA0CUZhA/kjmQiQMEldIb0aURn0gKXvxQQGwHHahUAkljl\n8j945BC8/XtFgJKha9BjVOxnAE5z8qNRsAl+nrtvnOvpseXd+r8XoYa1DozI2q3f\nuM3huumdewKBgQCdnEsRDvuPs1AVDleCyTpOR8DRb1/LlmDKNVWjiX73NmXQQ42i\nEUxQyyuGOUKb0K2rjAjblZ9hC0ZWLUxS5cWr+AD6Nm+OamdxjdrBLgsJIzi4glvR\n+XmLy4UdcKhVg1hfEgAxRnRybn95swiwajmrg8rSdChAhu3lsFi9nIKsHwKBgQCG\nGun4izi0eoBG5PLYW7Dk2cQf8tUyUs6r2wPv+0WwMmAkDbEsyRyilqlyaGiK41jM\nh9FgETJ6w3vcPKu7/XCZFbMkCzWq42fPj9NrEzcLNOlbeNVIe/Q9FabMYu8LKyn+\nZWkFAmW7ziP7WYZIFVlmiEb99Xdb2O2xhKqa8jofJwKBgQCcUSL6dsG5x4w/Vaci\ndgnKy4ECPnezvgx7wl9f3wk8IylKzUxz3gPjOGsO0X20exJSoO11STo+C+r2vyio\nd31o9ymTmDJ9bCbDW1g9HBr4Pc6bkmQ+SP9Kx2E3HJuo2Zji2joePAkJh1HVdkG7\nWcBLQg1Pn31c6LxK6hJowwa2vg==\n-----END PRIVATE KEY-----\n" json:"private_key"`
			ClientEmail             string `envconfig:"FIREBASE_CLIENT_EMAIL" default:"firebase-adminsdk-2k6lp@greendeco-2726b.iam.gserviceaccount.com" json:"client_email"`
			ClientID                string `envconfig:"FIREBASE_CLIENT_ID" default:"102751571694782723192" json:"client_id"`
			AuthURI                 string `envconfig:"FIREBASE_AUTH_URI" default:"https://accounts.google.com/o/oauth2/auth" json:"auth_uri"`
			TokenURI                string `envconfig:"FIREBASE_TOKEN_URI" default:"https://oauth2.googleapis.com/token" json:"token_uri"`
			AuthProviderX509CertURL string `envconfig:"FIREBASE_AUTH_PROVIDER_X509_CERT_URL" default:"https://www.googleapis.com/oauth2/v1/certs" json:"auth_provider_x509_cert_url"`
			ClientX509CertURL       string `envconfig:"CLIENT_X509_CERT_URL" default:"https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-2k6lp%40greendeco-2726b.iam.gserviceaccount.com" json:"client_x509_cert_url"`
			UniverseDomain          string `envconfig:"UNIVERSE_DOMAIN" default:"googleapis.com" json:"universe_domain"`
		}
		Bucket string `envconfig:"FIREBASE_BUCKET" default:"greendeco-2726b.appspot.com"`
	}
	VnPay struct {
		Secret     string `envconfig:"VNPAY_SECRET" default:"XNBCJFAKAZQSGTARRLGCHVZWCIOIGSHN"`
		TmnCode    string `envcofig:"VNPAY_TMNCODE" default:"CGXZLS0Z"`
		ReturnUrl  string `envconfig:"VNPAY_RETURN_URL" default:"http://localhost:8080/api/v1/payment/vnpay_return"`
		CancelUrl  string `envconfig:"VNPAY_CANCEL_URL" default:"http://localhost:3000/user/order"`
		SuccessUrl string `envconfig:"VNPAY_SUCCESS_URL" default:"http://localhost:3000/payment/received"`
		ErrorUrl   string `envconfig:"VNPAY_ERROR_URL" default:""`
	}
	PayPal struct {
		ClientId  string `envconfig:"PAYPAL_CLIENT" default:"ARd2n7YRPvm8jG2enCE7SXX5CLxlHSx0a8DlGfMmZEn18qKM8WHB8jHYoOputMIEMZLrhdZ1QO4bvsgD"`
		SecretKey string `envconfig:"PAYPAL_SECRET_KEY" default:"EOZ8Wy0m32AVqv1mCHQHs8FqSkitSiar1Dk1zLsFpLUaRAKzWXHVlFTTzwQ1jI79Svrbz1ljy4FKv_xq"`
		ReturnUrl string `envconfig:"PAYPAL_RETURN_URL" `
	}

	ExchangeMoneyApi struct {
		Url string `envconfig:"EXCHANGE_MONEY_API" default:"https://api.currencybeacon.com/v1/convert?api_key=R4XVKm6WBbcB13ACSc3SweC93YKmuoEi"`
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
