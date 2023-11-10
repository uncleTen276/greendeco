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
// @Summary create new cart item if cart have already existed update it's quantity by one
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
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
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

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": itemId,
	})
}

// @UpdateCartProduct() godoc
// @Summary cart item
// @Tags Cart
// @Param id path string true "id cart product"
// @Param todo body models.UpdateCartProduct true "cart"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/product/{id} [put]
// @Security Bearer
func UpdateCartProduct(c *fiber.Ctx) error {
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

	pId := c.Params("id")
	productId, err := uuid.Parse(pId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	updateProduct := &models.UpdateCartProduct{
		ID: productId,
	}
	if err := c.BodyParser(updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	cartItem, err := cartRepo.GetCartProductById(productId)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	if !isCartOwner(cartItem.Cart, *uid) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	if err := cartRepo.UpdateCartProductById(updateProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// @DeleteCartItemByCartId() godoc
// @Summary clear cart items
// @Tags Cart
// @Param id path string true "id cart"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/{id}/clear [delete]
// @Security Bearer
func DeleteCartItemByCartId(c *fiber.Ctx) error {
	id := c.Params("id")
	cartId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

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

	if !isCartOwner(cartId, *uid) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	if err := cartRepo.DeleteCartItemByCartId(cartId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// @DeleteCartItemById() godoc
// @Summary delete cart item
// @Tags Cart
// @Param id path string true "id cart"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/product/{id} [delete]
// @Security Bearer
func DeleteCartItemById(c *fiber.Ctx) error {
	pId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	cart, err := cartRepo.GetCartProductById(pId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

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

	if !isCartOwner(cart.Cart, *uid) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	if err := cartRepo.DeleteCartItemById(pId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// @GetCartById() godoc
// @Summary get cart by id
// @Tags Cart
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/{id} [Get]
// @Security Bearer
func GetCartById(c *fiber.Ctx) error {
	cId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

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

	if !isCartOwner(cId, *uid) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	cart, err := cartRepo.GetCartById(cId)
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
		Items:    cart,
		Page:     1,
		PageSize: 1,
		Next:     false,
		Prev:     false,
	})
}

// @GetCartProductById() godoc
// @Summary get cart product by id
// @Tags Cart
// @Param id path string true "id of cart product"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/product/{id} [Get]
// @Security Bearer
func GetCartProductById(c *fiber.Ctx) error {
	pId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

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

	cartRepo := repository.NewCartRepo(database.GetDB())
	cartProduct, err := cartRepo.GetCartProductById(pId)
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

	if !isCartOwner(cartProduct.Cart, *uid) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    cartProduct,
		Page:     1,
		PageSize: 1,
		Next:     false,
		Prev:     false,
	})
}

// @GetCartProductsByCartId() godoc
// @Summary query get cart products by cart Id
// @Description sort value can only asc or desc
// @Tags Cart
// @Param queries query models.BaseQuery false "default: limit = 10"
// @Param id path string true "id of cart"
// @Accept json
// @Produce json
// @Success 200 {object} models.BasePaginationResponse
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/{id}/product [Get]
// @Security Bearer
func GetCartProductsByCartId(c *fiber.Ctx) error {
	cId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	query := models.DefaultQuery()
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid filter",
		})
	}

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

	if !isCartOwner(cId, *uid) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	cartProductList, err := cartRepo.GetCartProductByCartId(cId, query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	nextPage := query.HaveNextPage(len(cartProductList))
	if nextPage {
		cartProductList = cartProductList[:len(cartProductList)-1]
	}

	return c.JSON(models.BasePaginationResponse{
		Items:    cartProductList,
		Page:     query.GetPageNumber(),
		PageSize: len(cartProductList),
		Next:     nextPage,
		Prev:     !query.IsFirstPage(),
	})
}

// @GetCartByOwnerId() godoc
// @Summary query cart by owner Id
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} models.BasePaginationResponse
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /cart/ [Get]
// @Security Bearer
func GetCartByOnwerId(c *fiber.Ctx) error {
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

	cartRepo := repository.NewCartRepo(database.GetDB())
	cart, err := cartRepo.GetCartByOwnerId(*uid)
	if err == models.ErrNotFound {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    cart,
		Page:     1,
		PageSize: 1,
		Next:     false,
		Prev:     false,
	})
}

func isCartOwner(cartId uuid.UUID, ownerId uuid.UUID) bool {
	cartRepo := repository.NewCartRepo(database.GetDB())
	cart, err := cartRepo.GetCartByOwnerId(ownerId)
	if err != nil {
		return false
	}

	return cart.ID == cartId
}
