package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
)

// @CreateColor() godoc
// @Summary create new color require admin permission
// @Tags Color
// @Param todo body models.CreateColor true "New color"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /color/ [post]
// @Security Bearer
func CreateColor(c *fiber.Ctx) error {
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

	newColor := &models.CreateColor{}
	if err := c.BodyParser(newColor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(newColor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	colorRepo := repository.NewColorRepo(database.GetDB())
	colorId, err := colorRepo.Create(newColor)
	if err != nil {
		if database.DetectDuplicateError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "record already exists",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": colorId,
	})
}

// @GetColorById() godoc
// @Summary get color by id
// @Tags Color
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} models.BasePaginationResponse
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /color/{id} [Get]
func GetColorById(c *fiber.Ctx) error {
	cId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	colorRepo := repository.NewColorRepo(database.GetDB())
	color, err := colorRepo.GetColorById(cId)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    color,
		PageSize: 1,
		Page:     1,
		Next:     false,
		Prev:     false,
	})
}

// @GetColors() godoc
// @Summary get colors
// @Tags Color
// @Accept json
// @Produce json
// @Success 200 {object} models.BasePaginationResponse
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /color/ [Get]
func GetColors(c *fiber.Ctx) error {
	colorRepo := repository.NewColorRepo(database.GetDB())
	colorList, err := colorRepo.All()
	if err == models.ErrNotFound {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    colorList,
		Page:     1,
		PageSize: len(colorList),
		Next:     false,
		Prev:     false,
	})
}

// @UpdateColorById() godoc
// @Summary update color
// @Tags Color
// @Param id path string true "id color"
// @Param todo body models.UpdateColor true "color"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /color/{id} [put]
// @Security Bearer
func UpdateColorById(c *fiber.Ctx) error {
	cId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	updateColor := &models.UpdateColor{
		ID: cId,
	}

	if err := c.BodyParser(updateColor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(updateColor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	colorRepo := repository.NewColorRepo(database.GetDB())
	if err := colorRepo.UpdateColor(updateColor); err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "soemthing bad happend :(",
		})

	}

	return c.SendStatus(fiber.StatusOK)
}
