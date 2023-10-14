package postgres

import (
	"context"
	"fmt"

	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/postgres/bridge"
	"github.com/eduardor2m/questao-certa/internal/app/entity/user"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
)

var _ repository.UserLoader = &UserPostgresRepository{}

type UserPostgresRepository struct {
	connectorManager
}

func (instance UserPostgresRepository) SignUp(u user.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	err = queries.SignUp(ctx, bridge.SignUpParams{
		ID:       u.ID(),
		Name:     u.Name(),
		Email:    u.Email(),
		Password: u.Password(),
		Admin:    false,
	})

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) SignIn(email string, password string) (*string, error) {
	return nil, nil
}

func (instance UserPostgresRepository) VerifyUserIsLoggedOrAdmin(token string) (*string, error) {
	return nil, nil
}

func NewUserPostgresRepository(connectorManager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{
		connectorManager: connectorManager,
	}
}
