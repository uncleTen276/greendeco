package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
)

// @GetRecommendProductsById() godoc
// @Summary get recomend products by Id
// @Tags Product
// @Param id path string true "id product"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /product/{id}/recommend/ [Get]
func GetRecommendProductsById(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	productRepo := repository.NewProductRepo(database.GetDB())
	recommends, err := productRepo.GetRecommendProducts(uid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    recommends,
		Page:     1,
		PageSize: len(recommends),
		Next:     false,
		Prev:     false,
	})
}

// @CreateRecommendProduct() godoc
// @Summary create new product recommend require admin permission
// @Tags Product
// @Param todo body models.CreateRecommend true "New recommend for product"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /product/recommend/ [post]
// @Security Bearer
func CreateRecommendProduct(c *fiber.Ctx) error {
	recommend := &models.CreateRecommend{}
	if err := c.BodyParser(recommend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(recommend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	if recommend.ProductId == recommend.RecommendId {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "can not recommend itself",
		})
	}

	productRepo := repository.NewProductRepo(database.GetDB())
	if err := productRepo.CreateRecommendProduct(recommend); err != nil {
		if database.DetectNotFoundContrainError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid recommend",
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
