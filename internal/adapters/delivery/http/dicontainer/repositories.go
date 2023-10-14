package dicontainer

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/mongodb"
	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/postgres"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
)

func GetQuestionRepository() repository.QuestionLoader {
	return mongodb.NewQuestionMongodbRepository(GetMongoConnectionManager())
}

func GetUserRepository() repository.UserLoader {
	return postgres.NewUserPostgresRepository(GetPsqlConnectionManager())
}

func GetMongoConnectionManager() *mongodb.DatabaseConnectorManager {
	return &mongodb.DatabaseConnectorManager{}
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectorManager {
	return &postgres.DatabaseConnectorManager{}
}
