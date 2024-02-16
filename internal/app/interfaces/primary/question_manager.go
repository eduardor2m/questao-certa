package primary

import (
	"mime/multipart"

	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
)

type QuestionManager interface {
	Create(questionReceived question.Question) error
	ImportByCSV(multipart.File) error
	List(page int) ([]question.Question, error)
	ListByFilter(filterReceived filter.Filter) ([]question.Question, error)
	DeleteByID(id string) error
	DeleteAll() error
}
