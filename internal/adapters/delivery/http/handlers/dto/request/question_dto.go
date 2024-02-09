package request

type QuestionDTO struct {
	Organization string   `json:"organization" example:"CESGRANRIO"`
	Model        string   `json:"model" example:"multiple_choice"`
	Year         string   `json:"year" example:"2023"`
	Discipline   string   `json:"discipline" example:"Engehnaria de Produção"`
	Topic        string   `json:"topic" example:"Administração da Produção"`
	Question     string   `json:"question" example:"Qual o objetivo da administração da produção?"`
	Options      []string `json:"options" example:"[\"Aumentar a produtividade\", \"Diminuir a produtividade\", \"Aumentar a qualidade\", \"Diminuir a qualidade\"]"`
	Answer       string   `json:"answer" example:"Aumentar a qualidade"`
}
