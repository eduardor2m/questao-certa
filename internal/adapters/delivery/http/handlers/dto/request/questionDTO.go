package request

import "github.com/google/uuid"

type MultipleChoiceDTO struct {
	ID           uuid.UUID `json:"id"`
	Organization string    `json:"organization"`
	Model        string    `json:"model"`
	Year         string    `json:"year"`
	Discipline   string    `json:"discipline"`
	Topic        string    `json:"topic"`
	Question     string    `json:"question"`
	Options      []string  `json:"options"`
	Answer       string    `json:"answer"`
}
