package dicontainer

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/mongodb"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
)

func GetQuestionRepository() repository.QuestionLoader {
	return mongodb.NewQuestionMongodbRepository(GetPsqlConnectionManager())
}

func GetPsqlConnectionManager() *mongodb.DatabaseConnectorManager {
	return &mongodb.DatabaseConnectorManager{}
}
