package mongodb_test

import (
	"testing"

	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/mongodb"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestQuestionMongodbRepository_CreateQuestion(t *testing.T) {
	mockManager := new(mongodb.MockConnectorManager)
	repo := mongodb.NewQuestionMongodbRepository(mockManager)

	mockDB := &mongo.Database{}
	mockManager.On("getConnection").Return(mockDB, nil)

	// mockDB.On("Collection", "questions").Return(mockDB.C("questions"))
	// mockDB.C("questions").On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)

	questionReceived, err := question.NewBuilder().WithQuestion("Question").WithAnswer("Answer").WithOptions([]string{"Option 1", "Option 2"}).Build()
	if err != nil {
		t.Error(err)
	}

	baseReceived, err := base.NewBuilder().WithOrganization("Organization").WithModel("Model").WithYear("2023").WithDiscipline("Discipline").WithTopic("Topic").Build()

	if err != nil {
		t.Error(err)
	}

	questionReceived.Base = *baseReceived
	err = repo.CreateQuestion(*questionReceived)
	assert.NoError(t, err)

	mockManager.AssertCalled(t, "getConnection")
	// mockDB.C("questions").AssertCalled(t, "InsertOne", mock.Anything, mock.Anything)

	// mockManager.AssertExpectations(t)
	// mockDB.C("questions").AssertExpectations(t)
}
