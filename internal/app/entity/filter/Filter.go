package filter

type Filter struct {
	organization string
	year         string
	topic        string
	discipline   string
	quantity     int64
}

func (instance *Filter) Organization() string {
	return instance.organization
}

func (instance *Filter) Year() string {
	return instance.year
}

func (instance *Filter) Topic() string {
	return instance.topic
}

func (instance *Filter) Discipline() string {
	return instance.discipline
}

func (instance *Filter) Quantity() int64 {
	return instance.quantity
}
