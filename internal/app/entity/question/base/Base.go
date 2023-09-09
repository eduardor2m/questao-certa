package base

import "github.com/google/uuid"

type Base struct {
	id           uuid.UUID // Identificador único da questão.
	organization string    // Organização ou banca que elaborou a questão.
	model        string    // Modelo da questão (por exemplo, múltipla escolha, dissertativa, etc.).
	year         string    // Ano em que a questão foi elaborada.
	discipline   string    // Conteúdo da questão.
	topic        string    // Tópico ou assunto ao qual a questão está relacionada.
}

func (instance *Base) ID() uuid.UUID {
	return instance.id
}

func (instance *Base) Organization() string {
	return instance.organization
}

func (instance *Base) Model() string {
	return instance.model
}

func (instance *Base) Year() string {
	return instance.year
}

func (instance *Base) Discipline() string {
	return instance.discipline
}

func (instance *Base) Topic() string {
	return instance.topic
}
