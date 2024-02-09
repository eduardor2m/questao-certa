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

func (instance UserPostgresRepository) SignUp(u user.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	err = queries.SignUp(ctx, bridge.SignUpParams{
		ID:        u.ID(),
		Name:      u.Name(),
		Email:     u.Email(),
		Password:  u.Password(),
		Admin:     false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) SignIn(email string, password string) (*string, error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	userDB, err := queries.SignIn(ctx, email)

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

func (instance UserPostgresRepository) VerifyUserIsLoggedOrAdmin(tokenJwt string) (*string, error) {
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

	userDB, err := queries.VerifyUserIsLoggedOrAdmin(ctx, userIdUUID)

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

func NewUserPostgresRepository(connectorManager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{
		connectorManager: connectorManager,
	}
}
