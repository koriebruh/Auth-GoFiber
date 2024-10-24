package handler

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/restful/api/model/web"
	"koriebruh/restful/api/service"
	"koriebruh/restful/api/utils"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuth(app *fiber.App, authService service.AuthService, authMid fiber.Handler) {
	h := authHandler{
		authService: authService,
	}

	app.Post("token/generate", h.GenerateToken)
	app.Get("token/validate", authMid, h.ValidateToken)
}

func (a authHandler) GenerateToken(ctx *fiber.Ctx) error {
	var request web.AuthRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(400)
	}

	token, err := a.authService.Authentication(ctx.Context(), request)
	if err != nil {
		return ctx.SendStatus(utils.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(token)
}

func (a authHandler) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	return ctx.Status(200).JSON(user)
}
