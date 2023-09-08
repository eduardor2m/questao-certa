package filter

type Filter struct {
	organization string
	year         string
	topic        string
	content      string
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

func (instance *Filter) Content() string {
	return instance.content
}
