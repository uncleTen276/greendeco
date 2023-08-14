package middlewares

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/pkg/configs"
)

func JWTProtected() func(c *fiber.Ctx) error {
	jwtwareConfig := jwtware.Config{
		SigningKey:     []byte(configs.AppConfig().Auth.JWTSecret),
		ContextKey:     "user",
		ErrorHandler:   jwtError,
		SuccessHandler: verifyTokenExpiration,
	}
	return jwtware.New(jwtwareConfig)
}

func verifyTokenExpiration(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	expires := int64(claims["exp"].(float64))
	if time.Now().Unix() > expires {
		return jwtError(c, errors.New("token expired"))
	}
	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
		Message: err.Error(),
	})
}
