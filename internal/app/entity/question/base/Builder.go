package base

import (
	"errors"
	"regexp"

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
	if model == "multiple_choice" || model == "true_or_false" {
		instance.Err = errors.New("Model is required and must be multiple_choice or true_or_false")
		return instance
	}
	instance.model = model
	return instance
}

func (instance *Builder) WithYear(year string) *Builder {
	if year == "" {
		instance.Err = errors.New("Year is required")
		return instance
	}

	regexp := regexp.MustCompile(`^[0-9]{4}$`)
	if !regexp.MatchString(year) {
		instance.Err = errors.New("Year must be a valid year")
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
