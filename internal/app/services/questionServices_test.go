package services

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/internal/app/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateQuestion(t *testing.T) {
	questionRepositoryMock := new(mocks.QuestionRepositoryMock)

	QuestionService := NewQuestionServices(questionRepositoryMock)

	questionRepositoryMock.On("CreateQuestion", mock.Anything).Return(nil)

	questionReceived, err := question.NewBuilder().WithQuestion("Question").WithAnswer("Answer").WithOptions([]string{"Option 1", "Option 2"}).Build()
	baseReceived, err := base.NewBuilder().WithOrganization("Organization").WithModel("Model").WithYear("Year").WithDiscipline("Discipline").WithTopic("Topic").Build()

	questionReceived.Base = *baseReceived

	err = QuestionService.CreateQuestion(*questionReceived)

	assert.Nil(t, err)
}
