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

// @GetOrderById() godoc
// @Summary GetOrderById() require owner or admin request
// @Tags Order
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /order/{id} [get]
// @Security Bearer
func GetOrderById(c *fiber.Ctx) error {
	oId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
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

	orderRepo := repository.NewOrderRepo(database.GetDB())
	order, err := orderRepo.GetOrderById(oId)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}
	}

	// if not admin or if not owner
	if order.OwnerId != *uid && !middlewares.GetAdminFromToken(token) {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    order,
		Page:     1,
		PageSize: 1,
		Prev:     false,
		Next:     false,
	})
}

// @GetOrderProductByOrderId() godoc
// @Summary GetOrderProductByOrderId() require owner or admin request
// @Tags Order
// @Param id path string true "id"
// @Param queries query models.BaseQuery false "default: limit = 10"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /order/{id}/product/ [get]
// @Security Bearer
func GetOrderProductByOrderId(c *fiber.Ctx) error {
	oId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
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

	// query
	query := models.DefaultQuery()
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid filter",
		})
	}

	orderRepo := repository.NewOrderRepo(database.GetDB())
	order, err := orderRepo.GetOrderById(oId)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}
	}

	if order.OwnerId != *uid && !middlewares.GetAdminFromToken(token) {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	products, err := orderRepo.GetOrderProductsByOrderId(oId, query)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happened :(",
		})
	}

	nextPage := query.HaveNextPage(len(products))
	if nextPage {
		products = products[:len(products)-1]
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    products,
		Page:     query.GetPageNumber(),
		PageSize: len(products),
		Next:     nextPage,
		Prev:     !query.IsFirstPage(),
	})
}

// @GetOrderByToken() godoc
// @Summary GetOrderByToken() require owner
// @Description "field" not working on swagger you can read models.ProductQueryField for fields query
// @Param queries query models.OrderQuery false "default: limit = 10"
// @Param fields query string false "fields query is json" example(field={"name":"hello"})
// @Tags Order
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /order/ [get]
// @Security Bearer
func GetOrderByToken(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	uid, err := middlewares.GetUserIdFromToken(token)
	query := &models.OrderQuery{
		BaseQuery: *models.DefaultQuery(),
	}

	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid filter",
		})
	}

	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	orderRepo := repository.NewOrderRepo(database.GetDB())
	query.Fields.OwnerId = uid
	orders, err := orderRepo.All(query)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}
	}

	nextPage := query.HaveNextPage(len(orders))
	if nextPage {
		orders = orders[:len(orders)-1]
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    orders,
		PageSize: len(orders),
		Page:     query.GetPageNumber(),
		Next:     nextPage,
		Prev:     !query.IsFirstPage(),
	})
}

// @UpdateOrderStatus() godoc
// @Summary UpdateOrderStatus() use to update order (status must follow the tree) draft -> processing -> completed -> cancelled
// @Tags Order
// @Param id path string true "order id"
// @Param todo body models.UpdateOrder true "product"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /order/{id} [put]
// @Security Bearer
func UpdateOrderStatus(c *fiber.Ctx) error {
	oId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	updateOrder := &models.UpdateOrder{
		OrderId: oId,
	}

	if err := c.BodyParser(updateOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validator := validators.NewValidator()
	if updateOrder.PaidAt == nil {
		if err := validator.StructExcept(updateOrder, "PaidAt"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid input found",
				Errors:  validators.ValidatorErrors(err),
			})
		}
	} else {
		if err := validator.Struct(updateOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid input found",
				Errors:  validators.ValidatorErrors(err),
			})
		}
	}

	if !validateOrderState(updateOrder) {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "order can not be updated",
		})
	}

	orderRepo := repository.NewOrderRepo(database.GetDB())
	if err := orderRepo.UpdateOrder(updateOrder); err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend",
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// @GetTotalOrderById() godoc
// @Summary GetTotalOrderById() use to get total for order
// @Tags Order
// @Param id path string true "order id"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /order/{id}/total [get]
// @Security Bearer
func GetTotalOrderById(c *fiber.Ctx) error {
	oId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	orderRepo := repository.NewOrderRepo(database.GetDB())
	order, err := orderRepo.GetOrderById(oId)
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

	total, err := orderRepo.GetTotalPaymentForOrder(oId)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid order",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	actualPrice := float64(total)
	if order.CouponDiscount != 0 {
		actualPrice *= float64(order.CouponDiscount) / 100
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total":        total,
		"actual_price": actualPrice,
	})
}

// @GetAllOrders() godoc
// @Summary GetAllOrders() require admin
// @Description "field" not working on swagger you can read models.ProductQueryField for fields query
// @Param queries query models.OrderQuery false "default: limit = 10"
// @Param fields query string false "fields query is json" example(field={"name":"hello"})
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} models.BasePaginationResponse
// @Failure 400,404,500 {object} models.ErrorResponse "Error"
// @Router /order/all/ [get]
// @Security Bearer
func GetAllOrders(c *fiber.Ctx) error {
	query := &models.OrderQuery{
		BaseQuery: *models.DefaultQuery(),
	}

	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid filter",
		})
	}

	orderRepository := repository.NewOrderRepo(database.GetDB())
	orders, err := orderRepository.All(query)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend",
		})

	}

	nextPage := query.HaveNextPage(len(orders))
	if nextPage {
		orders = orders[:len(orders)-1]
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    orders,
		PageSize: len(orders),
		Page:     query.GetPageNumber(),
		Next:     nextPage,
		Prev:     query.IsFirstPage(),
	})
}

// validateOrderState use to validate order before update
func validateOrderState(m *models.UpdateOrder) bool {
	orderRepo := repository.NewOrderRepo(database.GetDB())
	order, err := orderRepo.GetOrderById(m.OrderId)
	if err != nil {
		return false
	}

	if order.State == m.State {
		return true
	}

	childerState, ok := models.StatusTable[order.State]
	if !ok {
		return false
	}

	for _, state := range childerState {
		if state == m.State {
			return true
		}
	}

	return false
}
