package dtos

import "github.com/google/uuid"

type UserDB struct {
	ID       uuid.UUID `bson:"id"`
	Name     string    `bson:"name"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
	Admin    bool      `bson:"admin"`
}
