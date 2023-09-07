package multipleselect

import "github.com/eduardor2m/questao-certa/internal/app/entity/question/base"

type MultipleSelect struct {
	base.Base
	question string
	options  []string
	answer   []string
}
