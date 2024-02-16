package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	name       string
	email      string
	password   string
	admin      bool
	isActive   bool
	created_at time.Time
	updated_at time.Time
}

func (instance *User) ID() uuid.UUID {
	return instance.id
}

func (instance *User) Name() string {
	return instance.name
}

func (instance *User) Email() string {
	return instance.email
}

func (instance *User) Password() string {
	return instance.password
}

func (instance *User) Admin() bool {
	return instance.admin
}

func (instance *User) IsActive() bool {
	return instance.isActive
}

func (instance *User) CreatedAt() time.Time {
	return instance.created_at
}

func (instance *User) UpdatedAt() time.Time {
	return instance.updated_at
}
