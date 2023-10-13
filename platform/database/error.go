package database

import (
	"github.com/jackc/pgconn"
)

const (
	ErrorDBForeignKeyViolation = "23503"
	ErrorDBUnique              = "23505"
)

func DetectDuplicateError(err error) bool {
	errs, ok := err.(*pgconn.PgError)
	if !ok {
		return false
	}
	println(errs.Code)

	return errs.Code == ErrorDBUnique
}

func DetectNotFoundContrainError(err error) bool {
	errs, ok := err.(*pgconn.PgError)
	if !ok {
		return false
	}

	return errs.Code == ErrorDBForeignKeyViolation
}
