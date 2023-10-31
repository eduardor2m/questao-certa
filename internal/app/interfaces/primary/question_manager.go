package primary

import (
	"mime/multipart"

	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
)

type QuestionManager interface {
	CreateQuestion(questionReceived question.Question) error
	ImportQuestionsByCSV(multipart.File) error
	ListQuestions(page int) ([]question.Question, error)
	ListQuestionsByFilter(filterReceived filter.Filter) ([]question.Question, error)
	DeleteQuestion(id string) error
	DeleteAllQuestions() error
}
