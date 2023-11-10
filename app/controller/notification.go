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

// @CreateNotification() godoc
// @Summary create new notification return id
// @Tags Notification
// @Param todo body models.CreateNotification true "New notification"
// @Accept json
// @Produce		json
// @Success 200
// @Router /notification/ [post]
// @Security Bearer
func CreateNotification(c *fiber.Ctx) error {
	newNotification := &models.CreateNotification{}

	if err := c.BodyParser(newNotification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(newNotification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	notiRepo := repository.NewNotificationRepo(database.GetDB())
	notiId, err := notiRepo.Create(newNotification)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": notiId,
	})
}

// @SendNotification() godoc
// @Summary send notification to users
// @Tags Notification
// @Param todo body models.UserListNotification true "notification"
// @Accept json
// @Produce		json
// @Success 200
// @Router /notification/send [post]
// @Security Bearer
func SendNotification(c *fiber.Ctx) error {
	usersList := &models.UserListNotification{}
	if err := c.BodyParser(usersList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	notiRepo := repository.NewNotificationRepo(database.GetDB())
	if err := notiRepo.SendNotificationToUsers(usersList); err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "record not found",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// @GetNotificationByToken() godoc
// @Summary get notification by token
// @Tags Notification
// @Param queries query models.BaseQuery false "default: limit = 10"
// @Accept json
// @Produce		json
// @Success 200
// @Router /notification/ [get]
// @Security Bearer
func GetNotificationByToken(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "can't extract user info from request",
		})
	}

	query := models.DefaultQuery()
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	notiRepo := repository.NewNotificationRepo(database.GetDB())
	notifications, err := notiRepo.GetNotificationsByUserId(*userId, query)
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

	nextPage := query.HaveNextPage(len(notifications))
	if nextPage {
		notifications = notifications[:len(notifications)-1]
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    notifications,
		PageSize: len(notifications),
		Page:     query.GetPageNumber(),
		Next:     nextPage,
		Prev:     query.IsFirstPage(),
	})
}

// @UpdateReadNotification() godoc
// @Summary update state of notification (use for user => change unread to read)
// @Tags Notification
// @Param id path string true "id notification"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /notification/{id}/user [put]
// @Security Bearer
func UpdateReadNotification(c *fiber.Ctx) error {
	nId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	notiRepo := repository.NewNotificationRepo(database.GetDB())
	if err := notiRepo.UpdateReadNotification(nId); err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// @GetNotificationById() godoc
// @Summary get notification by id
// @Tags Notification
// @Param id path string true "id notification"
// @Accept json
// @Produce json
// @Success 200 {object} models.BasePaginationResponse
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /notification/{id} [get]
// @Security Bearer
func GetNotificationById(c *fiber.Ctx) error {
	nId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	notiRepo := repository.NewNotificationRepo(database.GetDB())
	noti, err := notiRepo.GetNotificationById(nId)
	if err != nil {
		if err == models.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Message: "record not found",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.BasePaginationResponse{
		Items:    noti,
		Page:     1,
		PageSize: 1,
		Next:     false,
		Prev:     false,
	})
}

// @UpdateNotification() godoc
// @Summary update notification info
// @Tags Notification
// @Param id path string true "id notification"
// @Param todo body models.UpdateNotification true "notification request"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /notification/{id} [put]
// @Security Bearer
func UpdateNotification(c *fiber.Ctx) error {
	nId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid id",
		})
	}

	updateNoti := &models.UpdateNotification{
		ID: nId,
	}
	if err := c.BodyParser(updateNoti); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
		})
	}

	validator := validators.NewValidator()
	if err := validator.Struct(updateNoti); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	notiRepo := repository.NewNotificationRepo(database.GetDB())
	_, err = notiRepo.GetNotificationById(nId)
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

	if err := notiRepo.UpdateNotificaionById(updateNoti); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happend :(",
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}
