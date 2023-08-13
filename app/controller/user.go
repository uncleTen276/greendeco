package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
)

// @CreateUser() godoc
// @Summary Create new user
// @Tag Auth
// @Param todo body models.CreateUser true "New User"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /auth/register [post]
func CreateUser(c *fiber.Ctx) error {
	user := &models.CreateUser{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": validators.ValidatorErrors(err),
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())
	// find user by identifier
	userExist, err := userRepo.FindUserByIdentifier(user.Identifier)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if userExist != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": "this user identifier already exists",
		})
	}

	// find user by email
	userExist, err = userRepo.FindUserByEmail(user.Email)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if userExist != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": "this user email already exists",
		})
	}

	err = userRepo.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return nil
}
