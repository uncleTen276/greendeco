package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/app/repository"
	"github.com/sekke276/greendeco.git/pkg/validators"
	"github.com/sekke276/greendeco.git/platform/database"
)

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

func sendNotitication(userId uuid.UUID, notiId uuid.UUID) error {
	notiRepo := repository.NewNotificationRepo(database.GetDB())
	if err := notiRepo.CreateNotificationUser(&models.CreateUserNotication{
		UserId:         userId,
		NotificationId: notiId,
		State:          "unread",
	}); err != nil {
		return err
	}

	return nil
}

func sendUserNotification(notification *models.CreateNotification, userId uuid.UUID) error {
	notiRepo := repository.NewNotificationRepo(database.GetDB())
	_, err := notiRepo.Create(notification)
	if err != nil {
		return err
	}
	return nil
}
