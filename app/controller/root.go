package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/models"
)

// GetPagination() function return pagination
// Default pagination offset is 1 and limit is 10
func GetPagination(c *fiber.Ctx) (*models.Pagination, error) {
	limitQr := c.Query("limit")
	offsetQr := c.Query("offset")

	pagination := models.DefaultPagination()

	if len(limitQr) > 0 {
		psInt, err := strconv.Atoi(limitQr)
		if err != nil {
			return nil, err
		}

		pagination.Limit = psInt
	}

	if len(offsetQr) > 0 {
		offsetInt, err := strconv.Atoi(offsetQr)
		if err != nil {
			return nil, err
		}

		pagination.OffSet = offsetInt
	}

	return pagination, nil
}
