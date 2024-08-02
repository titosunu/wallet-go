package utils

import (
	"errors"
	"net/http"

	"github.com/titosunu/wallet-go/core"
)

func GetHttpStatus(err error) int {
	if errors.Is(err, core.ErrAuthFailed) || errors.Is(err, core.ErrUserNotFound) {
			return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}