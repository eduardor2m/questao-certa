package mocks

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/stretchr/testify/mock"
)

type QuestionRepositoryMock struct {
	mock.Mock
}

func (instance *QuestionRepositoryMock) CreateQuestion(questionReceived question.Question) error {
	args := instance.Called(questionReceived)
	return args.Error(0)
}

func (instance *QuestionRepositoryMock) ImportQuestionsByCSV(questions []question.Question) error {
	args := instance.Called(questions)
	return args.Error(0)
}

func (instance *QuestionRepositoryMock) ListQuestions() ([]question.Question, error) {
	args := instance.Called()
	return args.Get(0).([]question.Question), args.Error(1)
}

func (instance *QuestionRepositoryMock) ListQuestionsByFilter(filterReceived filter.Filter) ([]question.Question, error) {
	args := instance.Called(filterReceived)
	return args.Get(0).([]question.Question), args.Error(1)
}

func (instance *QuestionRepositoryMock) DeleteQuestion(id string) error {
	args := instance.Called(id)
	return args.Error(0)
}

func (instance *QuestionRepositoryMock) DeleteAllQuestions() error {
	args := instance.Called()
	return args.Error(0)
}
