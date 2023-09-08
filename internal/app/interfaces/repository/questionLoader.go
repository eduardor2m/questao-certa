package repository

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
)

type QuestionLoader interface {
	CreateQuestion(Question multiplechoice.MultipleChoice) error
	ListQuestions() ([]multiplechoice.MultipleChoice, error)
	ListQuestionsByFilter(filter filter.Filter) ([]multiplechoice.MultipleChoice, error)
	DeleteQuestion(id string) error
	DeleteAllQuestions() error
}
