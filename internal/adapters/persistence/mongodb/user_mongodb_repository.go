package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/mongodb/utils/dtos"
	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/mongodb/utils/token"
	"github.com/eduardor2m/questao-certa/internal/app/entity/user"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

	defer instance.closeConnection(conn)

	ctx := context.Background()

	var userDB *dtos.UserDB

	err = conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).FindOne(ctx, bson.M{
		"email": userReceived.Email(),
	}).Decode(&userDB)
	if err != nil {
		return err
	}

	if userDB != nil {
		return fmt.Errorf("email already exists")
	}

	_, err = conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).InsertOne(ctx, bson.M{
		"id":       userReceived.ID(),
		"name":     userReceived.Name(),
		"email":    userReceived.Email(),
		"password": userReceived.Password(),
		"admin":    true,
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
	defer instance.closeConnection(conn)

	collection := conn.Collection(os.Getenv("MONGODB_COLLECTION_USER"))
	ctx := context.Background()

	var userDB *dtos.UserDB

	err = collection.FindOne(ctx, bson.M{"email": email}).Decode(&userDB)

	if err != nil {
		return nil, err
	}

	if userDB == nil {
		return nil, fmt.Errorf("user not found")
	}

	userFormatted, err := user.NewBuilder().WithID(userDB.ID).WithName(userDB.Name).WithEmail(userDB.Email).WithPassword(userDB.Password).WithAdmin(userDB.Admin).Build()
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFormatted.Password()), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password incorrect")
	}

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userFormatted.ID(),
		"authorized": true,
		"exp":        time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %v", err)
	}

	return &tokenString, nil
}

func (instance *UserMongodbRepository) VerifyUserIsLoggedOrAdmin(tokenReceived string) (*string, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	tokenJWT, err := token.StringToJWT(tokenReceived)
	if err != nil {
		return nil, fmt.Errorf("falha ao verificar token: %v", err)
	}

	claims, err := token.ExtractClainsFromJwtToken(tokenJWT)
	if err != nil {
		return nil, fmt.Errorf("falha ao verificar token: %v", err)
	}

	userIdFromToken := claims["user_id"]
	userAuthorizedFromToken := claims["authorized"].(bool)

	userIdFromTokenUUID, err := uuid.Parse(userIdFromToken.(string))
	if err != nil {
		return nil, fmt.Errorf("falha ao converter id para uuid: %v", err)
	}

	if userAuthorizedFromToken {
		var userDB *dtos.UserDB
		err := conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).FindOne(ctx, bson.M{
			"id": userIdFromTokenUUID,
		}).Decode(&userDB)

		if err != nil {
			return nil, fmt.Errorf("falha ao buscar usuário: %v", err)
		}

		userFormatted, err := user.NewBuilder().
			WithID(userDB.ID).
			WithName(userDB.Name).
			WithEmail(userDB.Email).
			WithPassword(userDB.Password).
			WithAdmin(userDB.Admin).Build()
		if err != nil {
			return nil, err
		}

		var userType string

		if userFormatted.Admin() {
			userType = "admin"
			return &userType, nil
		} else {
			userType = "user"
			return &userType, nil
		}
	}

	return nil, fmt.Errorf("user not authorized")
}

func NewUserMongodbRepository(connectorManager connectorManager) *UserMongodbRepository {
	return &UserMongodbRepository{connectorManager: connectorManager}
}
