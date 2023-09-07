package routes

import (
	_ "github.com/eduardor2m/questao-certa/internal/adapters/delivery/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Routes interface {
	Load(group *echo.Group)
}

type routes struct{}

func New() Routes {
	return &routes{}
}

func (instance *routes) Load(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)

	loadQuestionRoutes(group)
}
