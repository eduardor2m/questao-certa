package middlewares

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/labstack/echo/v4"
)

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		tokenAuthHeader := context.Request().Header.Get("Authorization")

		if tokenAuthHeader == "" {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		userServices := dicontainer.GetUserServices()

		_, err := userServices.CheckType(tokenAuthHeader[7:])

		if err != nil {
			return context.JSON(401, map[string]string{
				"message": "user not found: " + err.Error(),
			})
		}

		return next(context)
	}
}
