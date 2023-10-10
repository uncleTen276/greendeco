package database

import (
	"strings"

	"github.com/jackc/pgconn"
)

const (
	ErrorDBForeignKeyViolation = "23503"
)

func DetectDuplicateError(err error) bool {
	return strings.Contains(err.(*pgconn.PgError).Message, "duplicate key value violates unique constraint")
}

func DetectNotFoundContrainError(err error) bool {
	codeErr := err.(*pgconn.PgError).Code
	return codeErr == ErrorDBForeignKeyViolation
}
