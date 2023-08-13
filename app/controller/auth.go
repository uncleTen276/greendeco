package controller

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// IsValidPassword() use to compare hash password
func IsValidPassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

// GenerateToken() use to create accessToken and refreshToken
func GenerateTokens(uId uint) (*models.UserTokens, error) {
	auth := configs.AppConfig().Auth
	accessToken, err := GenerateAccessClaims(
		uId,
		time.Duration(auth.TokenExpire)*time.Minute,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateAccessClaims(
		uId,
		time.Duration(auth.RefreshTokenExpires)*time.Hour,
	)
	if err != nil {
		return nil, err
	}

	return &models.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GenerateAccessClaims() use to create new token
func GenerateAccessClaims(uId uint, timeNumber time.Duration) (string, error) {
	t := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uId
	claims["exp"] = t.Add(timeNumber).Unix()
	tokenString, err := token.SignedString([]byte(configs.AppConfig().Auth.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
