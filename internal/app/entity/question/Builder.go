package question

import "errors"

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
	if len(options) == 0 {
		instance.Err = errors.New("Options is required")
		return instance
	}
	instance.options = options
	return instance
}

func (instance *Builder) WithAnswer(answer string) *Builder {
	if answer == "" {
		instance.Err = errors.New("Answer is required")
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
