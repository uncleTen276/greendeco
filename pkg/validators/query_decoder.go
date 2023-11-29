package validators

import (
	"encoding/json"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/sekke276/greendeco.git/app/models"
)

// add decoder
func AddProductQueryDecoder() {
	fieldConverter := func(value string) reflect.Value {
		fieldType := models.ProductQueryField{}
		if err := json.Unmarshal([]byte(value), &fieldType); err != nil {
			return reflect.Value{}
		}
		return reflect.ValueOf(fieldType)
	}

	customField := fiber.ParserType{
		Customtype: models.ProductQueryField{},
		Converter:  fieldConverter,
	}

	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ParserType:        []fiber.ParserType{customField},
		ZeroEmpty:         true,
	})
}

func AddOrderQueryDecoder() {
	fieldConverter := func(value string) reflect.Value {
		fieldType := models.OrderQueryField{}
		if err := json.Unmarshal([]byte(value), &fieldType); err != nil {
			return reflect.Value{}
		}
		return reflect.ValueOf(fieldType)
	}

	customField := fiber.ParserType{
		Customtype: models.OrderQueryField{},
		Converter:  fieldConverter,
	}

	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ParserType:        []fiber.ParserType{customField},
		ZeroEmpty:         true,
	})
}
