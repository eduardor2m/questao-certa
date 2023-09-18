package primary

import "github.com/eduardor2m/questao-certa/internal/app/entity/user"

type UserManager interface {
	SignUp(user.User) error
	SignIn(email string, password string) (*string, error)
}
