package routes

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/middlewares"
	"github.com/labstack/echo/v4"
)

func loadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user", middlewares.Admin)
	userHandlers := dicontainer.GetUserHandler()

	userGroup.POST("", userHandlers.SignUp)
	userGroup.POST("/signin", userHandlers.SignIn)
	userGroup.GET("/verify", userHandlers.VerifyUserIsLoggedOrAdmin)
}
