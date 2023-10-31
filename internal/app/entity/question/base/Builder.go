package base

import (
	"errors"

	"github.com/eduardor2m/questao-certa/internal/app/utils/validator"
	"github.com/google/uuid"
)

type Builder struct {
	Base
	Err error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	if id == uuid.Nil {
		instance.Err = errors.New("ID is required")
		return instance
	}

	instance.id = id
	return instance
}

func (instance *Builder) WithOrganization(organization string) *Builder {
	if organization == "" {
		instance.Err = errors.New("Organization is required")
		return instance
	}

	instance.organization = organization
	return instance
}

func (instance *Builder) WithModel(model string) *Builder {
	valid, err := validator.IsModelValid(model)

	if !valid {
		instance.Err = err
	}

	instance.model = model
	return instance
}

func (instance *Builder) WithYear(year string) *Builder {
	valid, err := validator.IsYearValid(year)

	if !valid {
		instance.Err = err
		return instance
	}

	instance.year = year
	return instance
}

func (instance *Builder) WithDiscipline(discipline string) *Builder {
	if discipline == "" {
		instance.Err = errors.New("Discipline is required")
		return instance
	}

	instance.discipline = discipline
	return instance
}

func (instance *Builder) WithTopic(topic string) *Builder {
	if topic == "" {
		instance.Err = errors.New("Topic is required")
		return instance
	}
	instance.topic = topic
	return instance
}

func (instance *Builder) Build() (*Base, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}

	return &instance.Base, nil
}
