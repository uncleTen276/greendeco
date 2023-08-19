package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/app/templates"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/pkg/middlewares"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
	gomail "gopkg.in/mail.v2"
)

// @CreateUser() godoc
// @Summary Create new user
// @Tags Auth
// @Param todo body models.CreateUser true "New User"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /auth/register [post]
func CreateUser(c *fiber.Ctx) error {
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
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	if userExist != nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "this user identifier already exists",
		})
	}

	// find user by email
	userExist, err = userRepo.GetUserByEmail(user.Email)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	if userExist != nil {
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
	err = userRepo.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.Status(http.StatusCreated).SendString("create success")
}

// Login
// @Login godoc
// @Summary User Login
// @Description Use for login response the refresh_token and access_Token
// @Tags Auth
// @Accept json
// @Param todo body models.UserLogin true "Login"
// @Success 200 {object} models.UserTokens
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
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

	tokens, err := generateTokens(userExist)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
			Errors:  err.Error(),
		})
	}

	return c.JSON(tokens)
}

// GetUserInfo()
// @GetUserInfo godoc
// @Summary get user information by Id
// @Description route get user Id from token then get user information
// @Tags User
// @Accept json
// @Success 200 {object} models.User
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /user/me [get]
// @Security Bearer
func GetUserInfo(c *fiber.Ctx) error {
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
	user, err := userRepo.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	return c.JSON(user)
}

// ForgotPassword
// @ForgotPassword godoc
// @Summary option when user forgot password
// @Description send email to user for reset password
// @Tags Auth
// @Accept json
// @Param todo body controller.ForgotPassword.userEmailReq true "user email"
// @Success 200
// @Router /auth/forgot-password [post]
func ForgotPassword(c *fiber.Ctx) error {
	type userEmailReq struct {
		Email string `json:"email" validate:"required,email,lte=150"`
	}
	reqEmail := &userEmailReq{}
	if err := c.BodyParser(reqEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "Invalid request body",
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())

	user, err := userRepo.GetUserByEmail(reqEmail.Email)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	token, _ := generateTokens(user)

	go sendEmail(reqEmail.Email, token.AccessToken)
	return c.SendString("Please check your email")
}

// sendEmail use for send uemail to reset password
func sendEmail(email string, token string) {
	cfg := configs.AppConfig().SMTP
	newMessage := gomail.NewMessage()
	newMessage.SetHeader("From", "greendeco@gmail.com")
	newMessage.SetHeader("To", email)
	newMessage.SetHeader("Subject", "Reset Password")
	tmpl := template.Must(template.New("").Parse(templates.TemplateEmail))
	buff := new(bytes.Buffer)

	if err := tmpl.Execute(buff, models.EmailResponse{
		Link:   cfg.LinkResetPassword + fmt.Sprintf("/reset-password?token=%s", token),
		Sender: "greendeco@gmail.com",
	}); err != nil {
		fmt.Println("Failed to send email! Err: ", err)
	}

	newMessage.SetBody("text/html", buff.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 465, cfg.Email, cfg.Password)
	if err := dialer.DialAndSend(newMessage); err != nil {
		fmt.Println("Failed to send email! Err: ", err)
	}
}
