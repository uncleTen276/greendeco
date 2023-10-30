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

// @CreateOrderFromCart() godoc
// @Summary create new order from cart
// @Tags Order
// @Param todo body models.CreateCartOrder true "order request"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /order/ [post]
// @Security Bearer
func CreateOrderFromCart(c *fiber.Ctx) error {
	newOrderReq := &models.CreateCartOrder{}
	if err := c.BodyParser(newOrderReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())
	userId, err := middlewares.GetUserIdFromToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "can't extract user info from request",
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(newOrderReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	user, err := userRepo.GetUserById(*userId)
	if err != nil {
		if err != models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	newOrder := &models.Order{
		OwnerId:         user.ID,
		UserName:        user.FirstName + " " + user.LastName,
		UserEmail:       user.Email,
		UserPhoneNumber: user.PhoneNumber,
		State:           models.StatusDraft,
		ShippingAddress: newOrderReq.ShippingAddress,
	}

	if newOrderReq.CouponId.ID() != 0 {
		couponRepo := repository.NewCouponRepo(database.GetDB())
		coupon, err := couponRepo.GetCouponById(newOrderReq.CouponId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid input found",
			})
		}

		newOrder.Coupon = &coupon.ID
		newOrder.CouponDiscount = coupon.Discount
	}

	cartRepo := repository.NewCartRepo(database.GetDB())
	cartItem, err := cartRepo.GetAllCartProductByCartId(newOrderReq.CartId)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid cart",
			})
		}
	}

	if len(cartItem) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid cart",
		})
	}

	variantRepo := repository.NewVariantRepo(database.GetDB())
	orderItem := []*models.OrderProduct{}
	for _, item := range cartItem {
		variant, _ := variantRepo.FindById(item.Variant)
		newItem := &models.OrderProduct{
			VariantId:    item.Variant,
			VariantName:  variant.Name,
			VariantPrice: variant.Price,
			Quantity:     item.Quantity,
		}
		orderItem = append(orderItem, newItem)
	}

	orderRepo := repository.NewOrderRepo(database.GetDB())
	orderId, err := orderRepo.CreateOrderFromCart(newOrder, orderItem, newOrderReq.CartId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": orderId,
	})
}
