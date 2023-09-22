package question

import (
	"errors"

	"github.com/eduardor2m/questao-certa/internal/app/utils/validator"
)

type Builder struct {
	Question
	Err error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithQuestion(question string) *Builder {
	if question == "" {
		instance.Err = errors.New("Question is required")
		return instance
	}
	instance.question = question
	return instance
}

func (instance *Builder) WithOptions(options []string) *Builder {
	valid, err := validator.IsOptionsValid(options)

	if !valid {
		instance.Err = err
		return instance
	}

	instance.options = options
	return instance
}

func (instance *Builder) WithAnswer(answer string) *Builder {
	valid, err := validator.IsAnswerValid(answer)

	if !valid {
		instance.Err = err
		return instance
	}

	instance.answer = answer
	return instance
}

func (instance *Builder) Build() (*Question, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}

	return &instance.Question, nil
}
