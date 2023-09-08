package request

type FilterDTO struct {
	Organization string `json:"organization"`
	Year         string `json:"year"`
	Content      string `json:"content"`
	Topic        string `json:"topic"`
}
