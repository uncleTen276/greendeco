package controller

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plutov/paypal"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/platform/database"
)

// @CreateVnPayPayment() godoc
// @Summary create new order from cart
// @Tags Payment
// @Param todo body models.PaymentRequest true "order request"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /payment/vnpay_create [post]
// @Security Bearer
func CreateVnPayPayment(c *fiber.Ctx) error {
	newReq := &models.PaymentRequest{}
	if err := c.BodyParser(newReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	orderRepo := repository.NewOrderRepo(database.GetDB())
	order, err := orderRepo.GetOrderById(newReq.Id)
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

	url, err := createVNPayBill(order, c.IP())
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

	return c.JSON(fiber.Map{
		"callback_url": url,
	})
}

// @VnPay_Return() godoc
// @Summary vnPay return
// @Tags Payment
// @Param todo body models.PaymentRequest true "order request"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /payment/vnpay_return [Get]
func VnPay_Return(c *fiber.Ctx) error {
	vpnParams := c.Queries()
	secureHash := vpnParams["vnp_SecureHash"]
	delete(vpnParams, "vnp_SecureHash")
	delete(vpnParams, "vnp_SecureHashType")

	urlps := url.Values{}
	for k, v := range vpnParams {
		urlps.Add(k, v)
	}

	sortedParam := sortURLValues(urlps)
	mac := hmac.New(sha512.New, []byte(configs.AppConfig().VnPay.Secret))
	mac.Write([]byte(sortedParam.Encode()))
	sign := hex.EncodeToString(mac.Sum(nil))

	if secureHash == sign {
		orderId := vpnParams["vnp_TxnRef"]
		if vpnParams["vnp_ResponseCode"] == "24" {
			return c.Redirect(configs.AppConfig().VnPay.CancelUrl + "/" + orderId)
		}

		oId, err := uuid.Parse(orderId)
		// oId, err := uuid.Parse("078e2f28-d36f-467c-baf8-a91a0a878871")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid id",
			})
		}

		if vpnParams["vnp_ResponseCode"] != "00" {
			return c.Redirect(configs.AppConfig().VnPay.ErrorUrl)
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

		paidAt := time.Now().Format(time.RFC3339)
		if err := orderRepo.UpdateOrder(&models.UpdateOrder{
			OrderId:     order.ID,
			State:       order.State,
			PaidAt:      &paidAt,
			Description: order.Description,
		}); err != nil {
			return c.Redirect(configs.AppConfig().VnPay.ErrorUrl)
		}

		return c.Redirect(configs.AppConfig().VnPay.SuccessUrl) // error page
	} else {
		return c.Redirect(configs.AppConfig().VnPay.ErrorUrl)
	}
}

// @CreatePayPalPayment() godoc
// @Summary create payment with paypal
// @Tags Payment
// @Param todo body models.PaymentRequest true "request"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /payment/paypal_create [post]
// @Security Bearer
func CreatePayPalPayment(c *fiber.Ctx) error {
	newReq := &models.PaymentRequest{}
	if err := c.BodyParser(newReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	// find order
	orderRepo := repository.NewOrderRepo(database.GetDB())
	order, err := orderRepo.GetOrderById(newReq.Id)
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

	// get total
	total, err := orderRepo.GetTotalPaymentForOrder(order.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	actualPrice := float64(total)
	if order.CouponDiscount != 0 {
		actualPrice *= float64(order.CouponDiscount) / 100
	}

	totalString := fmt.Sprintf("%.0f", actualPrice)
	//
	cfg := configs.AppConfig().PayPal
	client, err := paypal.NewClient(cfg.ClientId, cfg.SecretKey, paypal.APIBaseSandBox)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	_, err = client.GetAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	nOrder, err := client.CreateOrder("CAPTURE", []paypal.PurchaseUnitRequest{
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value:    totalString,
			},
			ReferenceID: order.ID.String(),
			Description: "thanks",
		},
	}, &paypal.CreateOrderPayer{}, &paypal.ApplicationContext{
		BrandName: "greendeco",
		ReturnURL: "http://localhost:8080/api/v1",
	})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"order_id": nOrder.ID})
}

// @PayPalReturn() godoc
// @Summary paypal return
// @Tags Payment
// @Param todo body models.PayPalReturn true "request"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /payment/paypal_return [Post]
func PayPalReturn(c *fiber.Ctx) error {
	newReq := &models.PayPalReturn{}
	if err := c.BodyParser(newReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	cfg := configs.AppConfig().PayPal
	client, err := paypal.NewClient(cfg.ClientId, cfg.SecretKey, paypal.APIBaseSandBox)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	_, err = client.GetAccessToken()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "record not found",
			Errors:  err.Error(),
		})
	}

	_, err = client.CaptureOrder(newReq.ID, paypal.CaptureOrderRequest{})
	if err != nil {
		return err
	}

	paypalOrder, err := client.GetOrder(string(newReq.ID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	// update order
	orderRepo := repository.NewOrderRepo(database.GetDB())
	orderIdString, err := uuid.Parse(paypalOrder.PurchaseUnits[0].ReferenceID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	order, err := orderRepo.GetOrderById(orderIdString)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	paidAt := time.Now().Format(time.RFC3339)
	if err := orderRepo.UpdateOrder(&models.UpdateOrder{
		OrderId:     order.ID,
		State:       order.State,
		PaidAt:      &paidAt,
		Description: order.Description,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend",
		})
	}

	return c.Redirect(configs.AppConfig().PayPal.SuccessUrl)
}

func exchangeCurrencyFromUSDToVN(amount float64) (float64, error) {
	req := map[string]any{
		"from":   "USD",
		"to":     "VND",
		"amount": amount,
	}

	buf, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}

	agent := fiber.AcquireAgent()
	agent.Request().Header.SetMethod("GET")
	agent.Request().Header.SetContentType("application/json")
	agent.Request().SetRequestURI(configs.AppConfig().ExchangeMoneyApi.Url)
	agent.Body(buf)
	err = agent.Parse()
	if err != nil {
		return 0, err
	}
	statusCode, body, errs := agent.Bytes()
	if statusCode == fiber.StatusInternalServerError {
		return 0, errors.New("fail to get Currency")
	}

	if len(errs) > 0 {
		return 0, err
	}

	var response models.PaymentCurrenctResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	return math.Round(response.Value), nil
}

func createVNPayBill(order *models.Order, IP string) (string, error) {
	orderRepo := repository.NewOrderRepo(database.GetDB())
	total, err := orderRepo.GetTotalPaymentForOrder(order.ID)
	if err != nil {
		return "", err
	}

	actualPrice := float64(total)
	if order.CouponDiscount != 0 {
		actualPrice *= float64(order.CouponDiscount) / 100
	}

	totalVNDFloat, err := exchangeCurrencyFromUSDToVN(actualPrice)
	if err != nil {
		return "", err
	}

	totalString := fmt.Sprintf("%.0f", totalVNDFloat*100)
	cfgs := configs.AppConfig().VnPay
	t := time.Now().Format("20060102150405")
	v := url.Values{}
	v.Set("vnp_Version", "2.1.0")
	v.Set("vnp_Command", "pay")
	v.Set("vnp_TmnCode", cfgs.TmnCode)
	v.Set("vnp_Locale", "vn")
	v.Set("vnp_CurrCode", "VND")
	v.Set("vnp_TxnRef", order.ID.String()) // change bac
	v.Set("vnp_OrderInfo", "customer paid orderId")
	v.Set("vnp_OrderType", "other")
	v.Set("vnp_Amount", totalString)
	v.Set("vnp_ReturnUrl", cfgs.ReturnUrl)
	v.Set("vnp_IpAddr", IP)
	v.Set("vnp_CreateDate", t)

	sortedParam := sortURLValues(v)

	hash := hmac.New(sha512.New, []byte(cfgs.Secret))
	hash.Write([]byte(sortedParam.Encode()))

	sign := hex.EncodeToString(hash.Sum(nil))
	v.Set("vnp_SecureHash", sign)

	return models.VnPayUrl + v.Encode(), nil
}

func sortURLValues(values url.Values) url.Values {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	sortedValues := make(url.Values)
	for _, key := range keys {
		sortedValues[key] = values[key]
	}
	return sortedValues
}
