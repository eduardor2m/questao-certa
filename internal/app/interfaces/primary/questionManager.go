package primary

import (
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
)

type QuestionManager interface {
	CreateQuestion(question multiplechoice.MultipleChoice) error
}
