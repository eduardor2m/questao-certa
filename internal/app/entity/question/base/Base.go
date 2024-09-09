package base

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	id           uuid.UUID
	organization string
	model        string
	year         string
	discipline   string
	topic        string
	created_at   time.Time
	updated_at   time.Time
}

func (instance *Base) ID() uuid.UUID {
	return instance.id
}

func (instance *Base) Organization() string {
	return instance.organization
}

func (instance *Base) Model() string {
	return instance.model
}

func (instance *Base) Year() string {
	return instance.year
}

func (instance *Base) Discipline() string {
	return instance.discipline
}

func (instance *Base) Topic() string {
	return instance.topic
}

func (instance *Base) CreatedAt() time.Time {
	return instance.created_at
}

func (instance *Base) UpdatedAt() time.Time {
	return instance.updated_at
}
