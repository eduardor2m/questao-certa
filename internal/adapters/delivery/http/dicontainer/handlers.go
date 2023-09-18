package dicontainer

import "github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers"

func GetQuestionHandler() *handlers.QuestionHandler {
	return handlers.NewQuestionHandler(GetQuestionServices())
}

func GetUserHandler() *handlers.UserHandler {
	return handlers.NewUserHandler(GetUserServices())
}
