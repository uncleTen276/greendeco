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

// @CreateCategories() godoc
// @Summary create new category require admin permission
// @Tags Category
// @Param todo body models.CreateCategory true "New category"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /category/ [post]
// @Security Bearer
func CreateCategories(c *fiber.Ctx) error {
	newCategory := &models.CreateCategory{}

	if err := c.BodyParser(newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	categoryRepo := repository.NewCategoryRepository(database.GetDB())

	if err := categoryRepo.Create(newCategory); err != nil {
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

// @UpdateCategories() godoc
// @Summary update category by id
// @Tags Category
// @Param id path string true "id category update"
// @Param todo body models.UpdateCategory true "category"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /category/{id} [put]
// @Security Bearer
func UpdateCategories(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	newCategory := &models.UpdateCategory{}
	if !middlewares.GetAdminFromToken(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "you do have permission",
		})
	}

	id := c.Params("id")
	repo := repository.NewCategoryRepository(database.GetDB())
	category, err := repo.FindById(id)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	if category == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	if err := c.BodyParser(newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	newCategory.ID = category.ID

	if err := repo.UpdateById(newCategory); err != nil {
		if database.DetectDuplicateError(err) {
			return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
				Message: "record already exists",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// @DeleteCategories() godoc
// @Summary delete category by id
// @Tags Category
// @Param id path string true "id category update"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /category/{id} [delete]
// @Security Bearer
func DeleteCategories(c *fiber.Ctx) error {
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

	id := c.Params("id")
	repo := repository.NewCategoryRepository(database.GetDB())
	category, err := repo.FindById(id)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	if category == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	if err := repo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// GetAllCategory()
// @GetAllCategory godoc
// @Summary get categories
// @Tags Category
// @Param limit query int false "default: limit = 10"
// @Param offset query int false "default: offset = 1"
// @Accept json
// @Success 200 {array} models.BasePaginationResponse
// @Failure 400,500 {object} models.ErrorResponse "Error"
// @Router /category/ [get]
func GetAllCategory(c *fiber.Ctx) error {
	baseQuery := models.DefaultQuery()
	if err := c.QueryParser(baseQuery); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid filter",
		})
	}

	repo := repository.NewCategoryRepository(database.GetDB())
	pageOffset := baseQuery.Limit * (baseQuery.OffSet - 1)
	categories, err := repo.All(baseQuery.Limit+1, pageOffset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	nextPage := baseQuery.HaveNextPage(len(categories))
	if nextPage {
		categories = categories[:len(categories)-1]
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    categories,
		Page:     baseQuery.GetPageNumber(),
		PageSize: len(categories),
		Next:     nextPage,
		Prev:     !baseQuery.IsFirstPage(),
	})
}
