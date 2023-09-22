package repository

import "github.com/eduardor2m/questao-certa/internal/app/entity/user"

type UserLoader interface {
	SignUp(user.User) error
	SignIn(email string, password string) (*string, error)
	VerifyUserIsLoggedOrAdmin(token string) (*string, error)
}
