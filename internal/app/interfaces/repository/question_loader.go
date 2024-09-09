package repository

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
)

type QuestionLoader interface {
	Create(Question question.Question) error
	ImportByCSV(questions []question.Question) error
	List(page int) ([]question.Question, error)
	ListByFilter(filter filter.Filter) ([]question.Question, error)
	DeleteByID(id string) error
	DeleteAll() error
}
