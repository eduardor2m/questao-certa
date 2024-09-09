package primary

import "github.com/eduardor2m/questao-certa/internal/app/entity/user"

type UserManager interface {
	Register(user.User) error
	Authenticate(email string, password string) (*string, error)
	Delete(email string, username string) error
	List() ([]user.User, error)
	FindByEmail(email string) (*user.User, error)
	CheckType(token string) (*string, error)
}
