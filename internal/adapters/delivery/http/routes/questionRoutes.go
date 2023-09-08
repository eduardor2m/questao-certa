package routes

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/labstack/echo/v4"
)

func loadQuestionRoutes(group *echo.Group) {
	questionGroup := group.Group("/question")
	questionHandlers := dicontainer.GetQuestionHandler()

	questionGroup.POST("", questionHandlers.CreateQuestion)
	questionGroup.GET("", questionHandlers.ListQuestions)
	questionGroup.GET("/filter", questionHandlers.ListQuestionsByFilter)
	questionGroup.DELETE("/:id", questionHandlers.DeleteQuestion)
	questionGroup.DELETE("", questionHandlers.DeleteAllQuestions)
}
