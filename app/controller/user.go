package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

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
// @Summary create new user
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
			Message: "fail to create user",
		})
	}

	if userExist != nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "this user identifier already exists",
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
// @Description use for login response the access_Token
// @Tags Auth
// @Accept json
// @Param todo body models.UserLogin true "Login"
// @Success 200 {object} models.UserToken
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

	validate := validators.NewValidator()
	if err := validate.Struct(reqEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())

	user, err := userRepo.GetUserByEmail(reqEmail.Email)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	if err == models.ErrNotFound {
		return c.SendString("Please check your email")
	}

	config := configs.AppConfig().Auth
	token, err := generateToken(user, time.Duration(config.ShortTokenExpire*int(time.Minute)))
	if err != nil {
		println(err)
	}

	if err := sendEmail(reqEmail.Email, token); err != nil {
		println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happenned :(",
		})
	}

	return c.SendString("Please check your email")
}

// sendEmail use for send uemail to reset password
func sendEmail(email string, token string) error {
	cfg := configs.AppConfig().SMTP
	newMessage := gomail.NewMessage()
	newMessage.SetHeader("From", "greendeco@gmail.com")
	newMessage.SetHeader("To", email)
	newMessage.SetHeader("Subject", "Reset Password")
	tmpl := template.Must(template.New("").Parse(templates.TemplateEmail))
	buff := new(bytes.Buffer)

	if err := tmpl.Execute(buff, models.EmailResponse{
		Link: cfg.LinkResetPassword + fmt.Sprintf("/reset-password?token=%s", token),
	}); err != nil {
		fmt.Println("Failed to send email! Err: ", err)
	}

	newMessage.SetBody("text/html", buff.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 465, cfg.Email, cfg.Password)
	if err := dialer.DialAndSend(newMessage); err != nil {
		fmt.Println("Failed to send email! Err: ", err)
		return err
	}

	println("send email to", email)
	return nil
}

// UpdatePassword
// @Update godoc
// @Description Update Password
// @Tags Auth
// @Accept json
// @Param todo body controller.UpdatePassword.userPassword true "Updated Password"
// @Success 204
// @Security Bearer
// @Router /auth/password [put]
func UpdatePassword(c *fiber.Ctx) error {
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

	type userPassword struct {
		Password string `json:"password" validate:"required,lte=50,gte=8"`
	}

	req := &userPassword{}
	if err = c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	hashedPassword, err := generatePasswordHash([]byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())
	if err = userRepo.UpdatePasswordById(hashedPassword, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateUserInformation
// @Update godoc
// @Description update user information
// @Tags User
// @Accept json
// @Param todo body models.UpdateUser true "Updated UserInformation"
// @Success 200 {object} models.User
// @Security Bearer
// @Router /user/update [put]
func UpdateUserInformation(c *fiber.Ctx) error {
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

	userUpdate := &models.UpdateUser{}
	if err := c.BodyParser(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.ErrorResponse{
			Message: err.Error(),
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())

	userExist, err := userRepo.GetUserById(userId)
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

	userEmail, err := userRepo.GetUserByEmail(userUpdate.Email)
	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "fail to create user",
		})
	}

	if userEmail.ID != userExist.ID {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "this user email already exists",
		})
	}

	validate := validators.NewValidator()
	if userUpdate.Avatar != nil {
		if err := validate.Struct(userUpdate); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid input found",
				Errors:  validators.ValidatorErrors(err),
			})
		}
	} else {
		if err := validate.StructExcept(userUpdate, "avatar"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "invalid input found",
				Errors:  validators.ValidatorErrors(err),
			})
		}
	}

	if err := userRepo.UpdateUserInfor(userId, userUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&models.ErrorResponse{
			Message: "fail to update user",
		})
	}

	userRes, err := userRepo.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(userRes)
}
