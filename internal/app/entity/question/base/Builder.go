package base

import "github.com/google/uuid"

type Builder struct {
	Base
	Err error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	instance.id = id
	return instance
}

func (instance *Builder) WithOrganization(organization string) *Builder {
	instance.organization = organization
	return instance
}

func (instance *Builder) WithModel(model string) *Builder {
	instance.model = model
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

func (instance *Builder) Build() (*Base, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}

	return &instance.Base, nil
}
