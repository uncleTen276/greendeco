package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
)

// @CreateReview() godoc
// @Summary create new review for product
// @Tags Review
// @Param todo body models.CreateReview true "New review"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /review/ [post]
// @Security Bearer
func CreateReview(c *fiber.Ctx) error {
	newReview := &models.CreateReview{}
	if err := c.BodyParser(newReview); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(newReview); err != nil {
		return err
	}

	reviewRepo := repository.NewReviewRepo(database.GetDB())
	if err := reviewRepo.Create(newReview); err != nil {
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

// @GetReviewById() godoc
// @Summary get review by id
// @Tags Review
// @Param id path string true "review id"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /review/{id} [get]
func GetReviewById(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	reviewRepo := repository.NewReviewRepo(database.GetDB())
	review, err := reviewRepo.FindById(uid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    review,
		Page:     1,
		PageSize: 1,
		Next:     false,
		Prev:     false,
	})
}

// @GetReviewById() godoc
// @Summary get review by product id
// @Tags Review
// @Param queries query models.ReviewQuery false "default: limit = 10"
// @Param id path string true "product id"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /review/product/{id} [get]
func GetReviewByProductId(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}
	query := &models.ReviewQuery{
		BaseQuery: *models.DefaultQuery(),
	}

	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	reviewRepo := repository.NewReviewRepo(database.GetDB())
	reviews, err := reviewRepo.FindReviewsByProductId(&uid, query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	nextPage := query.HaveNextPage()
	if nextPage {
		reviews = reviews[:len(reviews)-1]
	}

	return c.JSON(models.BasePaginationResponse{
		Items:    reviews,
		Page:     query.GetPageNumber(),
		PageSize: len(reviews),
		Next:     nextPage,
		Prev:     !query.IsFirstPage(),
	})
}
