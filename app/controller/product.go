package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
)

// @CreateProduct() godoc
// @Summary Create new product require admin permission
// @Tags Product
// @Param todo body models.CreateProduct true "New product"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /product/ [post]
// @Security Bearer
func CreateProduct(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	if !middlewares.GetAdminFromToken(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "you do have permission",
		})
	}

	newProduct := &models.CreateProduct{}
	if err := c.BodyParser(newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	productRepo := repository.NewProductRepository(database.GetDB())
	if err := productRepo.Create(newProduct); err != nil {
		if database.DetectNotFoundContrainError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid category",
			})
		}

		if database.DetectDuplicateError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "record already exists",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.Status(fiber.StatusCreated).SendString("create success")
}
