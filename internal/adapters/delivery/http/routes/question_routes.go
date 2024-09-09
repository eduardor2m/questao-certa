package routes

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/middlewares"
	"github.com/labstack/echo/v4"
)

func loadQuestionRoutes(group *echo.Group) {
	questionGroup := group.Group("/question")
	questionHandlers := dicontainer.GetQuestionHandler()

	questionGroup.POST("", middlewares.Admin(questionHandlers.Create))
	questionGroup.POST("/import", middlewares.Admin(questionHandlers.ImportByCSV))
	questionGroup.GET("/:page", middlewares.Admin(questionHandlers.List))
	questionGroup.POST("/filter", middlewares.GuardMiddleware(questionHandlers.ListByFilter))
	questionGroup.DELETE("/:id", middlewares.Admin(questionHandlers.DeleteByID))
	questionGroup.DELETE("", middlewares.Admin(questionHandlers.DeleteAll))
}
