package repository

import "github.com/eduardor2m/questao-certa/internal/app/entity/user"

type UserLoader interface {
	SignUp(user.User) error
	SignIn(email string, password string) (*string, error)
	DeleteUserTest(email string, username string) error
	ListUsers() ([]user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	VerifyUserIsLoggedOrAdmin(token string) (*string, error)
}
