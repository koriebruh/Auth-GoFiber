package middleware

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/restful/api/service"
	"koriebruh/restful/api/utils"
	"strings"
)

func Authentication(authService service.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(401)
		}

		validate, err := authService.Validate(ctx.Context(), token)
		if err != nil {
			return ctx.SendStatus(utils.GetHttpStatus(err))
		}

		ctx.Locals("x-user", validate)
		return ctx.Next()
	}
}
