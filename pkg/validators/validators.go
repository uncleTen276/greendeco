package validators

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateUUID, uuid.UUID{})
	return validate
}

func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		errMsg := fmt.Sprintf("validation failed on '%s' tag", err.Tag())
		param := err.Param()
		if param != "" {
			errMsg = fmt.Sprintf("%s. allow: %s", errMsg, param)
		}
		fields[err.Field()] = errMsg
	}

	return fields
}

func ValidateUUID(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		return valuer.String()
	}
	return nil
}
