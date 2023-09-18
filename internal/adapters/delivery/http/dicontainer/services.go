package dicontainer

import (
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/eduardor2m/questao-certa/internal/app/services"
)

func GetQuestionServices() primary.QuestionManager {
	return services.NewQuestionServices(GetQuestionRepository())
}

func GetUserServices() primary.UserManager {
	return services.NewUserServices(GetUserRepository())
}
