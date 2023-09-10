package question

type Builder struct {
	Question
	Err error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithQuestion(question string) *Builder {
	instance.question = question
	return instance
}

func (instance *Builder) WithOptions(options []string) *Builder {
	instance.options = options
	return instance
}

func (instance *Builder) WithAnswer(answer string) *Builder {
	instance.answer = answer
	return instance
}

func (instance *Builder) Build() (*Question, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}

	return &instance.Question, nil
}
