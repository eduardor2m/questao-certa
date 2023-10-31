package repository

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
)

type QuestionLoader interface {
	CreateQuestion(Question question.Question) error
	ImportQuestionsByCSV(questions []question.Question) error
	ListQuestions(page int) ([]question.Question, error)
	ListQuestionsByFilter(filter filter.Filter) ([]question.Question, error)
	DeleteQuestion(id string) error
	DeleteAllQuestions() error
}
