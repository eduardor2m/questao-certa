package request

import "github.com/google/uuid"

type QuestionDTO struct {
	ID           uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Organization string    `json:"organization" example:"CESGRANRIO"`
	Model        string    `json:"model" example:"multiple_choice"`
	Year         string    `json:"year" example:"2019"`
	Discipline   string    `json:"discipline" example:"ENGENHARIA DE PRODUÇÃO"`
	Topic        string    `json:"topic" example:"ADMINISTRAÇÃO DA PRODUÇÃO"`
	Question     string    `json:"question" example:"Qual o objetivo da administração da produção?"`
	Options      []string  `json:"options" example:"[\"Aumentar a produtividade\", \"Diminuir a produtividade\", \"Aumentar a qualidade\", \"Diminuir a qualidade\"]"`
	Answer       string    `json:"answer" example:"Aumentar a qualidade"`
}
