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

// @CreateVariant() godoc
// @Summary Create new variant require admin permission
// @Tags Variant
// @Param todo body models.CreateVariant true "New variant"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /variant/ [post]
// @Security Bearer
func CreateVariant(c *fiber.Ctx) error {
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

	newVariant := &models.CreateVariant{}
	if err := c.BodyParser(newVariant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(newVariant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	variantRepo := repository.NewVariantRepo(database.GetDB())

	if err := variantRepo.Create(newVariant); err != nil {
		if database.DetectDuplicateError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "record already exists",
			})
		}

		if database.DetectNotFoundContrainError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid product",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.Status(fiber.StatusCreated).SendString("create success")
}

// @GetVariantsByProductId() godoc
// @Summary get variants from product
// @Tags Variant
// @Param id path string true "product id"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /variant/product/{id} [get]
func GetVariantsByProductId(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}
	repo := repository.NewVariantRepo(database.GetDB())
	variants, err := repo.GetVariantsByProductId(uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "id not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    variants,
		PageSize: len(variants),
		Page:     1,
	})
}
