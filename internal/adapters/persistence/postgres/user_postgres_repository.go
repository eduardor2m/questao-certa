package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/postgres/bridge"
	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/utils/token"
	"github.com/eduardor2m/questao-certa/internal/app/entity/user"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var _ repository.UserLoader = &UserPostgresRepository{}

type UserPostgresRepository struct {
	connectorManager
}

func (instance UserPostgresRepository) Register(u user.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	timeNow := time.Now()

	err = queries.Register(ctx, bridge.RegisterParams{
		ID:        u.ID(),
		Name:      u.Name(),
		Email:     u.Email(),
		Password:  u.Password(),
		Admin:     true,
		IsActive:  true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	})

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) Authenticate(email string, password string) (*string, error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	userDB, err := queries.Authenticate(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("error comparing password: %v", err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userDB.ID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return nil, fmt.Errorf("error generating token: %v", err)
	}

	return &tokenString, nil

}

func (instance UserPostgresRepository) List() ([]user.User, error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	usersDB, err := queries.List(ctx)

	if err != nil {
		return nil, fmt.Errorf("error getting users: %v", err)
	}

	users := make([]user.User, 0)

	for _, userDB := range usersDB {
		userFormatted, err := user.NewBuilder().WithID(userDB.ID).WithIsActive(userDB.IsActive).WithCreatedAt(userDB.CreatedAt).WithUpdatedAt(userDB.UpdatedAt).WithAdmin(userDB.Admin).WithName(userDB.Name).WithEmail(userDB.Email).WithPassword(userDB.Password).Build()
		if err != nil {
			return nil, fmt.Errorf("error formatting user: %v", err)
		}

		users = append(users, *userFormatted)
	}

	return users, nil
}

func (instance UserPostgresRepository) FindByEmail(email string) (*user.User, error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	userDB, err := queries.FindByEmail(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	userFormatted, err := user.NewBuilder().WithID(userDB.ID).WithCreatedAt(userDB.CreatedAt).WithUpdatedAt(userDB.UpdatedAt).WithAdmin(userDB.Admin).WithName(userDB.Name).WithEmail(userDB.Email).WithPassword(userDB.Password).Build()
	if err != nil {
		return nil, fmt.Errorf("error formatting user: %v", err)
	}

	return userFormatted, nil
}

func (instance UserPostgresRepository) CheckType(tokenJwt string) (*string, error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	tokenReceived, err := token.StringToJWT(tokenJwt)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	claims, err := token.ExtractClainsFromJwtToken(tokenReceived)
	if err != nil {
		return nil, fmt.Errorf("error extracting claims from token: %v", err)
	}

	userID := claims["user_id"].(string)

	userIdUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("error parsing user id: %v", err)
	}

	userDB, err := queries.CheckType(ctx, userIdUUID)

	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	var typeUser string

	if userDB.Admin {
		typeUser = "admin"
	} else {
		typeUser = "user"
	}

	return &typeUser, nil
}

func (instance UserPostgresRepository) Delete(email string, name string) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	userDB, err := queries.FindByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	switch {
	case userDB.Name != name:
		return fmt.Errorf("error deleting user: name does not match")
	case userDB.Email != email:
		return fmt.Errorf("error deleting user: email does not match")
	case !userDB.IsActive:
		return fmt.Errorf("error deleting user: user is not active")
	case userDB.Admin:
		return fmt.Errorf("error deleting user: user is admin")
	}

	err = queries.Delete(ctx,
		bridge.DeleteParams{
			Email: email,
			Name:  name,
		},
	)

	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}

func NewUserPostgresRepository(connectorManager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{
		connectorManager: connectorManager,
	}
}
