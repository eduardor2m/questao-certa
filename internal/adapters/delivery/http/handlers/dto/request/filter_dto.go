package request

type FilterDTO struct {
	Organization string `json:"organization" example:"CESGRANRIO"`
	Year         string `json:"year" example:"2023"`
	Discipline   string `json:"discipline" example:"Engenharia de Produção"`
	Topic        string `json:"topic" example:"Administração da Produção"`
	Quantity     int64  `json:"quantity" example:"10"`
}
