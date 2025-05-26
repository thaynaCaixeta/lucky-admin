package service

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func CreateNewGameFailure() error {
	return huma.NewError(
		http.StatusInternalServerError,
		"create_new_game_failed",
		fmt.Errorf("failed to create new game"),
	)
}
