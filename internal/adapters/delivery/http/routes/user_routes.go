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
	userGroup.POST("/signin", userHandlers.Authenticate)
	userGroup.GET("/verify", middlewares.Admin(userHandlers.CheckType))
	userGroup.GET("/list", middlewares.Admin(userHandlers.List))
	userGroup.GET("/:email", middlewares.Admin(userHandlers.FindByEmail))
	userGroup.DELETE("/delete", middlewares.Admin(userHandlers.Delete))
}
