package controller

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"golang.org/x/crypto/bcrypt"
)

func generatePasswordHash(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// IsValidPassword() use to compare hash password
func isValidPassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

// GenerateToken() use to create accessToken and refreshToken
func generateTokens(uId string) (*models.UserTokens, error) {
	auth := configs.AppConfig().Auth
	accessToken, err := generateAccessClaims(
		uId,
		time.Duration(auth.TokenExpire)*time.Minute,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateAccessClaims(
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
func generateAccessClaims(uId string, timeNumber time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": uId,
		"exp": time.Now().Add(timeNumber).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(configs.AppConfig().Auth.JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
