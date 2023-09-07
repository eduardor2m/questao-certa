package repository

import (
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
)

type QuestionLoader interface {
	CreateQuestion(Question multiplechoice.MultipleChoice) error
}
