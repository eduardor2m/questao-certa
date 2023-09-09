package multiplechoice

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
)

type MultipleChoice struct {
	base.Base
	question string   // Enunciado da questão.
	options  []string // Opções de resposta.
	answer   string   // Resposta correta.
}

func (instance *MultipleChoice) Question() string {
	return instance.question
}

func (instance *MultipleChoice) Options() []string {
	return instance.options
}

func (instance *MultipleChoice) Answer() string {
	return instance.answer
}
