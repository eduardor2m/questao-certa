package user

import (
	"time"

	"github.com/eduardor2m/questao-certa/internal/app/utils/validator"
	"github.com/google/uuid"
)

type Builder struct {
	User
	Err error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	instance.id = id
	return instance
}

func (instance *Builder) WithName(name string) *Builder {
	instance.name = name
	return instance
}

func (instance *Builder) WithEmail(email string) *Builder {
	valid, err := validator.IsEmailValid(email)

	if err != nil {
		instance.Err = err
		return instance
	}

	if !valid {
		instance.Err = err
		return instance
	}

	instance.email = email
	return instance
}

func (instance *Builder) WithPassword(password string) *Builder {
	valid, err := validator.IsPasswordValid(password)

	if err != nil {
		instance.Err = err
		return instance
	}

	if !valid {
		instance.Err = err
		return instance
	}

	instance.password = password
	return instance
}

func (instance *Builder) WithAdmin(admin bool) *Builder {
	instance.admin = admin
	return instance
}

func (instance *Builder) WithIsActive(isActive bool) *Builder {
	instance.isActive = isActive
	return instance
}

func (instance *Builder) WithCreatedAt(created_at time.Time) *Builder {
	instance.created_at = created_at
	return instance
}

func (instance *Builder) WithUpdatedAt(updated_at time.Time) *Builder {
	instance.updated_at = updated_at
	return instance
}

func (instance *Builder) Build() (*User, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}

	return &instance.User, nil
}
