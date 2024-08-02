package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/titosunu/wallet-go/core"
	"github.com/titosunu/wallet-go/infrastructure/utils"
	"github.com/titosunu/wallet-go/payloads"
)

type authApi struct {
	userService core.UserService
}

func NewAuth(app *fiber.App, userService core.UserService, authMid fiber.Handler) {
	handler := authApi{
		userService: userService,
	}

	app.Post("token/generate", handler.GenerateToken)
	app.Get("token/validate", authMid, handler.ValidateToken)
}

func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req payloads.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		statusCode := utils.GetHttpStatus(err)
    return ctx.SendStatus(statusCode)
	}

  return ctx.Status(fiber.StatusOK).JSON(token)
}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	return ctx.Status(200).JSON(user)
}
