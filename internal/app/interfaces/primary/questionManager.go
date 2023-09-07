package primary

import (
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
)

type QuestionManager interface {
	CreateQuestion(question multiplechoice.MultipleChoice) error
	ListQuestions() ([]multiplechoice.MultipleChoice, error)
	ListQuestionsByOrganization(organization string) ([]multiplechoice.MultipleChoice, error)
	DeleteQuestion(id string) error
	DeleteAllQuestions() error
}
