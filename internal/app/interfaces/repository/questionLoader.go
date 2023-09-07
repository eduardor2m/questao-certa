package repository

import (
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
)

type QuestionLoader interface {
	CreateQuestion(Question multiplechoice.MultipleChoice) error
	ListQuestions() ([]multiplechoice.MultipleChoice, error)
	ListQuestionsByOrganization(organization string) ([]multiplechoice.MultipleChoice, error)
	DeleteQuestion(id string) error
	DeleteAllQuestions() error
}
