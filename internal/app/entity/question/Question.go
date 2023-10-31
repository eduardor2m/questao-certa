package question

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
)

type Question struct {
	base.Base
	question string
	options  []string
	answer   string
}

func (instance *Question) Question() string {
	return instance.question
}

func (instance *Question) Options() []string {
	return instance.options
}

func (instance *Question) Answer() string {
	return instance.answer
}
