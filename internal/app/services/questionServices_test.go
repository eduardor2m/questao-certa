package services

import (
	"testing"

	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/internal/app/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateQuestion(t *testing.T) {
	questionRepositoryMock := new(mocks.QuestionRepositoryMock)

	QuestionService := NewQuestionServices(questionRepositoryMock)

	questionRepositoryMock.On("CreateQuestion", mock.Anything).Return(nil)

	questionReceived, err := question.NewBuilder().WithQuestion("Question").WithAnswer("Answer").WithOptions([]string{"Option 1", "Option 2"}).Build()
	if err != nil {
		t.Error(err)
	}
	baseReceived, err := base.NewBuilder().WithOrganization("Organization").WithModel("Model").WithYear("2023").WithDiscipline("Discipline").WithTopic("Topic").Build()
	if err != nil {
		t.Error(err)
	}
	questionReceived.Base = *baseReceived

	err = QuestionService.CreateQuestion(*questionReceived)

	assert.Nil(t, err)
}

func TestListQuestions(t *testing.T) {
	questionRepositoryMock := new(mocks.QuestionRepositoryMock)

	QuestionService := NewQuestionServices(questionRepositoryMock)

	questionRepositoryMock.On("ListQuestions", mock.Anything).Return([]question.Question{}, nil)

	_, err := QuestionService.ListQuestions(10)

	assert.Nil(t, err)
}
