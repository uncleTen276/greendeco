package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Port     string `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" default:"postgres"`
		Password string `envconfig:"DB_PASSWORD" default:"postgres"`
		Name     string `envConfig:"DB_NAME" default:"greendeco"`
		SSLMode  string `envConfig:"DB_SSL_MODE" default:"disable"`
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
