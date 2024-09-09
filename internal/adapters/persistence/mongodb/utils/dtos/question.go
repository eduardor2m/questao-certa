package dtos

import (
	"time"

	"github.com/google/uuid"
)

type QuestionDB struct {
	ID           uuid.UUID `bson:"id"`
	Organization string    `bson:"organization"`
	Model        string    `bson:"model"`
	Year         string    `bson:"year"`
	Discipline   string    `bson:"discipline"`
	Topic        string    `bson:"topic"`
	Question     string    `bson:"question"`
	Answer       string    `bson:"answer"`
	Options      []string  `bson:"options"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}
