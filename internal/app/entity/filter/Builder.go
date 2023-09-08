package filter

type Builder struct {
	Filter
	Err error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithOrganization(organization string) *Builder {
	instance.organization = organization
	return instance
}

func (instance *Builder) WithYear(year string) *Builder {
	instance.year = year
	return instance
}

func (instance *Builder) WithContent(content string) *Builder {
	instance.content = content
	return instance
}

func (instance *Builder) WithTopic(topic string) *Builder {
	instance.topic = topic
	return instance
}

func (instance *Builder) Build() (*Filter, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}

	return &instance.Filter, nil
}
