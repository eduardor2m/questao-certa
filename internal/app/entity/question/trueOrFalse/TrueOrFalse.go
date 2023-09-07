package trueorfalse

import "github.com/eduardor2m/questao-certa/internal/app/entity/question/base"

type TrueOrFalse struct {
	base.Base
	Question string // Enunciado da questão.
	Answer   bool   // Resposta correta.
}
