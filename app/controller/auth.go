package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/platform/database"
	"golang.org/x/crypto/bcrypt"
)

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Auth
// @Produce json
// @Param request body controller.RefreshToken.request true "refresh_token"
// @Success 200 {object} models.UserTokens
// @Failure 400,401,403,404,500 {object} models.ErrorResponse "Error"
// @Router /auth/refresh-token [post]
func RefreshToken(c *fiber.Ctx) error {
	type request struct {
		Refresh_Token string `json:"refresh_token"`
	}

	tokenString := &request{}
	if err := c.BodyParser(tokenString); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	uId, err := parseToken(tokenString.Refresh_Token)
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(models.ErrorResponse{
			Message: "invalid or missing token",
		})
	}
	userRepo := repository.NewUserRepo(database.GetDB())
	user, err := userRepo.GetUserById(uId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "Id not found",
		})
	}

	tokens, err := generateTokens(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
			Errors:  err,
		})
	}

	return c.JSON(models.UserTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokenString.Refresh_Token,
	})
}

// use to parse Token return user id
func parseToken(tokenString string) (string, error) {
	var userId string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return configs.AppConfig().Auth.JWTSecret, nil
	})
	if err != nil {
		return userId, fiber.NewError(fiber.StatusUnauthorized, "invalid or missing token")
	}

	if !token.Valid {
		return userId, fiber.NewError(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userId, fiber.NewError(fiber.StatusUnauthorized)
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return userId, fiber.NewError(fiber.StatusUnauthorized, "token is out of date")
	}

	userId = claims["user_id"].(string)
	return userId, nil
}

// generatePasswordHash() use to generate hashed password
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

// generateToken() use to create accessToken and refreshToken
func generateTokens(user *models.User) (*models.UserTokens, error) {
	auth := configs.AppConfig().Auth
	accessToken, err := generateAccessClaims(
		user,
		time.Duration(auth.TokenExpire)*time.Minute,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateAccessClaims(
		user,
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
func generateAccessClaims(user *models.User, timeNumber time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"admin":   user.IsAdmin,
		"exp":     time.Now().Add(timeNumber).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(configs.AppConfig().Auth.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
