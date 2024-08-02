package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/titosunu/wallet-go/core"
	"github.com/titosunu/wallet-go/infrastructure/utils"
)

func Authenticate(userService core.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(401)
		}
		user, err := userService.ValidateToken(ctx.Context(), token)
		if err != nil {
			return ctx.SendStatus(utils.GetHttpStatus(err))
		}
		ctx.Locals("x-user", user)
		return ctx.Next()
	}
}