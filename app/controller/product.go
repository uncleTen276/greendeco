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

	productRepo := repository.NewProductRepo(database.GetDB())
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

// @UpdateProduct() godoc
// @Summary update product require admin permission
// @Tags Product
// @Param id path string true "id product update"
// @Param todo body models.UpdateProduct true "product"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /product/{id}/update [post]
// @Security Bearer
func UpdateProduct(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	if !middlewares.GetAdminFromToken(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "you don't have permission",
		})
	}

	updateProduct := &models.UpdateProduct{}
	if err := c.BodyParser(updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}
	updateProduct.ID = c.Params("id")

	validate := validators.NewValidator()
	if err := validate.Struct(updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	productRepo := repository.NewProductRepo(database.GetDB())
	if err := productRepo.UpdateById(updateProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// @DeleteProduct() godoc
// @Summary delete product by id require admin permission
// @Tags Product
// @Param id path string true "id product"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /product/{id}/delete [delete]
// @Security Bearer
func DeleteProduct(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	if !middlewares.GetAdminFromToken(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "you don't have permission",
		})
	}

	id := c.Params("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	productRepo := repository.NewProductRepo(database.GetDB())
	if err := productRepo.Delete(uuid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// @DeleteProduct() godoc
// @Summary query get products
// @Description "field" not working on swagger you can read models.ProductQueryField for fields query
// @Tags Product
// @Param queries query models.ProductQuery false "default: limit = 10"
// @Param fields query string false "fields query is json" example(field={"name":"hello"})
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /product/ [Get]
func GetProducts(c *fiber.Ctx) error {
	query := &models.ProductQuery{
		BaseQuery: *models.DefaultQuery(),
	}

	err := c.QueryParser(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	productRepo := repository.NewProductRepo(database.GetDB())
	products, err := productRepo.All(*query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	if query.HaveNextPage() {
		products = products[:len(products)-1]
	}

	return c.JSON(models.BasePaginationResponse{
		Items:    products,
		Page:     (query.BaseQuery.OffSet + query.Limit - 1) / query.Limit,
		PageSize: len(products),
		Next:     query.HaveNextPage(products),
		Prev:     !query.IsFirstPage(),
	})
}
