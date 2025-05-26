package repository

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func NewDatabaseError(detail string, err error) error {
	return huma.NewError(
		http.StatusInternalServerError,
		"database_error",
		fmt.Errorf("%s: %w", detail, err),
	)
}

func NewAdminNotFoundError(username string) error {
	return huma.NewError(
		http.StatusNotFound,
		"admin_not_found",
		fmt.Errorf("admin with username '%s' not found", username),
	)
}

func NewUUIDGenerationError(err error) error {
	return huma.NewError(
		http.StatusInternalServerError,
		"uuid_generation_failed",
		err,
	)
}

func NewTransactionCommitError(err error) error {
	return huma.NewError(
		http.StatusInternalServerError,
		"transaction_commit_failed",
		err,
	)
}
