package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Name     string `envconfig:"DB_NAME" default:"greendeco"`
		Port     string `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" default:"postgres"`
		Password string `envconfig:"DB_PASSWORD" default:"postgres"`
		SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
	}
	Auth struct {
		JWTRefreshToken     string `envconfig:"JWT_REFRESH_SECRET" default:"refresh-secret"`
		JWTSecret           string `envconfig:"JWT_SECRET" default:"token-secret"`
		TokenExpire         int    `envconfig:"TOKEN_EXPIRE" default:"15"`
		RefreshTokenExpires int    `envconfig:"REFRESH_TOKEN_EXPIRE" default:"720"`
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
