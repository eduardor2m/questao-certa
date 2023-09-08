package primary

import (
	"mime/multipart"

	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
)

type QuestionManager interface {
	CreateQuestion(question multiplechoice.MultipleChoice) error
	ImportQuestionsByCSV(multipart.File) error
	ListQuestions() ([]multiplechoice.MultipleChoice, error)
	ListQuestionsByFilter(filter filter.Filter) ([]multiplechoice.MultipleChoice, error)
	DeleteQuestion(id string) error
	DeleteAllQuestions() error
}
