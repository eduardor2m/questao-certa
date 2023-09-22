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
	questionGroup.POST("/import", questionHandlers.ImportQuestionsByCSV)
	questionGroup.GET("/:page", questionHandlers.ListQuestions)
	questionGroup.POST("/filter", questionHandlers.ListQuestionsByFilter)
	questionGroup.DELETE("/:id", questionHandlers.DeleteQuestion)
	questionGroup.DELETE("", questionHandlers.DeleteAllQuestions)
}
