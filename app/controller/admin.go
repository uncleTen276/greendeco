package controller

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
)

// GetAllCustomers()
// @GetUserInfo godoc
// @Summary get customers
// @Description route get customer for admin
// @Tags Admin
// @Accept json
// @Success 200 {array} models.User
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /admin/cusotmers [get]
// @Security Bearer
func GetAllCustomers(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(token)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())

	userExist, err := userRepo.GetUserById(*userId)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	if userExist == nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	if !userExist.IsAdmin {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(models.ErrorResponse{
			Message: "don't have permission",
		})
	}

	adminRepo := repository.NewAdminRepo(database.GetDB())
	users, err := adminRepo.GetCustomer(0, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// LoginForAdmin
// @LoginForAdmin godoc
// @Summary User Login
// @Description use for login response the access_Token
// @Tags Admin
// @Accept json
// @Param todo body models.UserLogin true "Login"
// @Success 200 {object} models.UserToken
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /admin/login/ [post]
func LoginForAdmin(c *fiber.Ctx) error {
	user := &models.UserLogin{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())
	userExist, err := userRepo.GetUserByIdentifier(user.Identifier)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&models.ErrorResponse{
			Message: "user not found",
		})
	}

	if !isValidPassword(userExist.Password, user.Password) {
		return c.Status(fiber.StatusForbidden).JSON(&models.ErrorResponse{
			Message: "wrong password",
		})
	}

	if !userExist.IsAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(&models.ErrorResponse{
			Message: "you dont have permission",
		})
	}

	config := configs.AppConfig().Auth
	token, err := generateToken(userExist, time.Duration(config.TokenExpire*int(time.Minute)*24))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
			Errors:  err.Error(),
		})
	}

	return c.JSON(models.UserToken{
		AccessToken: token,
	})
}

// @CreateAdminAccount() godoc
// @Summary create new user
// @Tags Auth
// @Param todo body models.CreateUser true "New User"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /admin/register [post]
func CreateAdminAccount(c *fiber.Ctx) error {
	user := &models.CreateUser{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())
	// find user by identifier
	userExist, err := userRepo.GetUserByIdentifier(user.Identifier)
	if userExist != nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "this user identifier already exists",
		})
	}

	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "fail to create user",
		})
	}

	// find user by email
	userEmail, err := userRepo.GetUserByEmail(user.Email)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "fail to create user",
		})
	}

	if userEmail != nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "this user email already exists",
		})
	}

	hashedPassword, err := generatePasswordHash([]byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	user.Password = hashedPassword
	err = userRepo.CreateForStaff(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.Status(http.StatusCreated).SendString("create success")
}
