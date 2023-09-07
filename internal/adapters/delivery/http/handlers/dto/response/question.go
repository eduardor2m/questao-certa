package response

import "github.com/google/uuid"

type MultipleChoice struct {
	ID           uuid.UUID `json:"id"`
	Organization string    `json:"organization"`
	Model        string    `json:"model"`
	Year         string    `json:"year"`
	Content      string    `json:"content"`
	Topic        string    `json:"topic"`
	Question     string    `json:"question"`
	Answer       string    `json:"answer"`
	Options      []string  `json:"options"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type InfoResponse struct {
	Message string `json:"message"`
}
