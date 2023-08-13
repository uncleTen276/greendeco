package models

import "errors"

var ErrNotFound = errors.New("object not found")

type ErrorResponse struct {
	Message string `json:"msg"`
	Errors  any    `json:"errors,omitempty"`
}
