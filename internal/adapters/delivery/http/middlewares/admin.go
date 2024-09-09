package middlewares

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/labstack/echo/v4"
)

func Admin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		tokenAuthHeader := context.Request().Header.Get("Authorization")

		if tokenAuthHeader == "" {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		userServices := dicontainer.GetUserServices()

		userType, err := userServices.CheckType(tokenAuthHeader[7:])

		if err != nil {
			return context.JSON(401, map[string]string{
				"message": "error verifying user: " + err.Error(),
			})
		}

		if *userType != "admin" {
			return context.JSON(401, map[string]string{
				"message": "user unauthorized, only admin can access this resource",
			})
		}

		return next(context)
	}
}
