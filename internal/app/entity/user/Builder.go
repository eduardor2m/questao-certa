package user

import "github.com/google/uuid"

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
	instance.email = email
	return instance
}

func (instance *Builder) WithPassword(password string) *Builder {
	instance.password = password
	return instance
}

func (instance *Builder) WithAdmin(admin bool) *Builder {
	instance.admin = admin
	return instance
}

func (instance *Builder) Build() (*User, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}

	return &instance.User, nil
}
