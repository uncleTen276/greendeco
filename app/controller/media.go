package controller

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"github.com/sekke276/greendeco.git/platform/storage"
)

// UploadImage() use to upload image to file storage using firebase storage
func UploadImage(ctx context.Context, fileInput []byte, fileName string, token string) error {
	stor, err := storage.GetStorage(ctx)
	if err != nil {
		return err
	}

	bucket, err := stor.DefaultBucket()
	if err != nil {
		return err
	}

	writer := bucket.Object(fileName).NewWriter(ctx)
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": token}
	if _, err := io.Copy(writer, bytes.NewReader(fileInput)); err != nil {
		return err
	}

	defer writer.Close()

	return nil
}

// @PostMedia() godoc
// @Summary create new image return image
// @Tags Media
// @ID	image
// @Produce		json
// @Accept	multipart/form-data
// @Param image formData file true "upfile"
// @Success 200
// @Router /media/upload [post]
// @Security Bearer
func PostMedia(c *fiber.Ctx) error {
	ctx := context.Background()
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
			Errors:  err.Error(),
		})
	}

	buffer, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "file not found",
			Errors:  err.Error(),
		})
	}
	defer buffer.Close()

	content, err := io.ReadAll(buffer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "file not found",
			Errors:  err.Error(),
		})
	}

	name := strings.Split(file.Filename, ".")
	token := uuid.New().String()
	fileName := token + "." + name[1]
	if err := UploadImage(ctx, content, fileName, token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
			Errors:  err.Error(),
		})
	}

	cfg := configs.AppConfig()
	link := "https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s"
	imageLink := fmt.Sprintf(link, cfg.Storage.Bucket, fileName, token)

	return c.SendString(imageLink)
}
