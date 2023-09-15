package request

type FilterDTO struct {
	Organization string `json:"organization" example:"CESGRANRIO"`
	Year         string `json:"year" example:"2019"`
	Discipline   string `json:"discipline" example:"ENGENHARIA DE PRODUÇÃO"`
	Topic        string `json:"topic" example:"ADMINISTRAÇÃO DA PRODUÇÃO"`
	Quantity     int64  `json:"quantity" example:"10"`
}
