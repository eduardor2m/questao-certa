package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/app/entity/user"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var _ repository.UserLoader = &UserMongodbRepository{}

type UserMongodbRepository struct {
	connectorManager
}

func (instance *UserMongodbRepository) SignUp(userReceived user.User) error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	ctx := context.Background()

	var userDTO request.UserDTO

	err = conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).FindOne(ctx, bson.M{
		"email": userReceived.Email(),
	}).Decode(&userDTO)

	if err == nil {
		return fmt.Errorf("email already exists")
	}

	_, err = conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).InsertOne(ctx, bson.M{
		"id":       userReceived.ID(),
		"name":     userReceived.Name(),
		"email":    userReceived.Email(),
		"password": userReceived.Password(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (instance *UserMongodbRepository) SignIn(email string, password string) (*string, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	var user request.UserDTO

	err = conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	userPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("falha ao comparar senha: %v", err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return nil, fmt.Errorf("falha ao criar token: %v", err)
	}

	return &tokenString, nil
}

func NewUserMongodbRepository(connectorManager connectorManager) *UserMongodbRepository {
	return &UserMongodbRepository{connectorManager: connectorManager}
}
