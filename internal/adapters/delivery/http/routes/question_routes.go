package routes

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/middlewares"
	"github.com/labstack/echo/v4"
)

func loadQuestionRoutes(group *echo.Group) {
	questionGroup := group.Group("/question")
	questionHandlers := dicontainer.GetQuestionHandler()

	questionGroup.POST("", middlewares.Admin(questionHandlers.CreateQuestion))
	questionGroup.POST("/import", middlewares.Admin(questionHandlers.ImportQuestionsByCSV))
	questionGroup.GET("/:page", middlewares.Admin(questionHandlers.ListQuestions))
	questionGroup.POST("/filter", middlewares.GuardMiddleware(questionHandlers.ListQuestionsByFilter))
	questionGroup.DELETE("/:id", middlewares.Admin(questionHandlers.DeleteQuestion))
	questionGroup.DELETE("", middlewares.Admin(questionHandlers.DeleteAllQuestions))
}
