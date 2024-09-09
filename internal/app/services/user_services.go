package services

import (
	"github.com/eduardor2m/questao-certa/internal/app/entity/user"
	"golang.org/x/crypto/bcrypt"

	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"github.com/google/uuid"
)

var _ primary.UserManager = (*UserServices)(nil)

type UserServices struct {
	userRepository repository.UserLoader
}

func (instance *UserServices) Register(userReceiced user.User) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReceiced.Password()), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	userFormatted, err := user.NewBuilder().WithID(id).WithName(userReceiced.Name()).WithPassword(string(passwordHash)).WithEmail(userReceiced.Email()).Build()

	if err != nil {
		return err
	}

	return instance.userRepository.Register(*userFormatted)
}

func (instance *UserServices) Authenticate(email string, password string) (*string, error) {
	return instance.userRepository.Authenticate(email, password)
}

func (instance *UserServices) Delete(email string, username string) error {
	return instance.userRepository.Delete(email, username)
}

func (instance *UserServices) List() ([]user.User, error) {
	return instance.userRepository.List()
}

func (instance *UserServices) FindByEmail(email string) (*user.User, error) {
	return instance.userRepository.FindByEmail(email)
}

func (instance *UserServices) CheckType(token string) (*string, error) {
	return instance.userRepository.CheckType(token)
}

func NewUserServices(userRepository repository.UserLoader) *UserServices {
	return &UserServices{userRepository: userRepository}
}
