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

// @CreateCart() godoc
// @Summary create new cart
// @Tags Cart
// @Param todo body models.CreateCart true "New cart"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/ [post]
// @Security Bearer
func CreateCart(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	uid, err := middlewares.GetUserIdFromToken(token)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	newCart := &models.CreateCart{
		Owner: *uid,
	}
	if err := c.BodyParser(newCart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(newCart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	cartId, err := cartRepo.Create(newCart)
	if err != nil {
		if database.DetectNotFoundContrainError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid user",
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": cartId,
	})
}

// @CreateCartProduct() godoc
// @Summary create new cart item
// @Tags Cart
// @Param todo body models.CreateCartProduct true "New cart"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/product/ [post]
// @Security Bearer
func CreateCartProduct(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	uid, err := middlewares.GetUserIdFromToken(token)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	newItem := &models.CreateCartProduct{}
	if err := c.BodyParser(newItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(newItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	if !isCartOwner(newItem.Cart, *uid) {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	if !validators.ValidateActiveVariant(newItem.Variant) {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid product",
		})
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	itemId, err := cartRepo.CreateCartProduct(newItem)
	if err != nil {
		if database.DetectNotFoundContrainError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid input",
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": itemId,
	})
}

func isCartOwner(cartId uuid.UUID, ownerId uuid.UUID) bool {
	cartRepo := repository.NewCartRepo(database.GetDB())
	cart, err := cartRepo.GetCartByOwnerId(ownerId)
	if err != nil {
		return false
	}

	return cart.Owner == ownerId
}
