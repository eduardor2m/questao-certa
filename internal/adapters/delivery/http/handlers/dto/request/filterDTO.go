package request

type FilterDTO struct {
	Organization string `json:"organization"`
	Year         string `json:"year"`
	Discipline   string `json:"discipline"`
	Topic        string `json:"topic"`
	Quantity     int64  `json:"quantity"`
}
