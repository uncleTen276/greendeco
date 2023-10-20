package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
	"github.com/sekke276/greendeco.git/platform/database"
)

// GetAllCustomers()
// @GetUserInfo godoc
// @Summary get customers
// @Description route get customer for admin
// @Tags Admin
// @Accept json
// @Success 200 {array} models.User
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /admin/cusotmers [get]
// @Security Bearer
func GetAllCustomers(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(token)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())

	userExist, err := userRepo.GetUserById(*userId)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	if userExist == nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	if !userExist.IsAdmin {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(models.ErrorResponse{
			Message: "don't have permission",
		})
	}

	adminRepo := repository.NewAdminRepo(database.GetDB())
	users, err := adminRepo.GetCustomer(0, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
