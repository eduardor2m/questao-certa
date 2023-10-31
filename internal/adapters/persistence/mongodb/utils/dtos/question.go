package dtos

import "github.com/google/uuid"

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
}
