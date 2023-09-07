package services

import (
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
)

var _ primary.QuestionManager = (*QuestionServices)(nil)

type QuestionServices struct {
	questionRepository repository.QuestionLoader
}

func (instance *QuestionServices) CreateQuestion(question multiplechoice.MultipleChoice) error {
	return instance.questionRepository.CreateQuestion(question)
}

func (instance *QuestionServices) ListQuestions() ([]multiplechoice.MultipleChoice, error) {
	return instance.questionRepository.ListQuestions()
}

func (instance *QuestionServices) ListQuestionsByOrganization(organization string) ([]multiplechoice.MultipleChoice, error) {
	return instance.questionRepository.ListQuestionsByOrganization(organization)
}

func (instance *QuestionServices) DeleteQuestion(id string) error {
	return instance.questionRepository.DeleteQuestion(id)
}

func (instance *QuestionServices) DeleteAllQuestions() error {
	return instance.questionRepository.DeleteAllQuestions()
}

func NewQuestionServices(questionRepository repository.QuestionLoader) *QuestionServices {
	return &QuestionServices{questionRepository: questionRepository}
}
