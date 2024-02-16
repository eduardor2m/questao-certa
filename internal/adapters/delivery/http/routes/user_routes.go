package routes

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/middlewares"
	"github.com/labstack/echo/v4"
)

func loadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")

	userHandlers := dicontainer.GetUserHandler()

	userGroup.POST("", userHandlers.Register)
	userGroup.POST("/authenticate", userHandlers.Authenticate)
	userGroup.GET("/checktype", middlewares.GuardMiddleware(userHandlers.CheckType))
	userGroup.GET("", middlewares.Admin(userHandlers.List))
	userGroup.GET("/:email", middlewares.Admin(userHandlers.FindByEmail))
	userGroup.DELETE("", middlewares.Admin(userHandlers.Delete))
}
