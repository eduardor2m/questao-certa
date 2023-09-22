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
		"admin":    false,
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

	userRow, err := conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).FindOne(ctx, bson.M{
		"email": email,
	}).DecodeBytes()

	if err != nil {
		return nil, err
	}

	//Call of bsoncore.Value.StringValue on binary type

	//Call of bsoncore.Value.StringValue on binary type

	subtipy, data := userRow.Lookup("id").Binary()

	if subtipy != 0 {
		return nil, fmt.Errorf("falha ao converter id para uuid: %v", err)
	}

	userDBId, err := uuid.FromBytes(data)

	fmt.Println(userDBId)

	if err != nil {
		return nil, fmt.Errorf("falha ao converter id para uuid: %v", err)
	}

	userDB, err := user.NewBuilder().WithID(userDBId).WithName(userRow.Lookup("name").StringValue()).WithPassword(userRow.Lookup("password").StringValue()).WithEmail(userRow.Lookup("email").StringValue()).Build()

	if err != nil {
		return nil, err
	}

	fmt.Println(userDB)

	userPassword := userDB.Password()

	if err != nil {
		return nil, fmt.Errorf("falha ao converter id para uuid: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("falha ao comparar senha: %v", err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userDB.ID()
	claims["authorized"] = true

	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return nil, fmt.Errorf("falha ao criar token: %v", err)
	}

	return &tokenString, nil
}

func (instance *UserMongodbRepository) VerifyUserIsLoggedOrAdmin(tokenReceived string) (*string, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenReceived, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("falha ao verificar token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFromToken := claims["user_id"]
		userAuthorizedFromToken := claims["authorized"].(bool)
		userIdFromTokenUUID, err := uuid.Parse(userIdFromToken.(string))

		if err != nil {
			return nil, fmt.Errorf("falha ao converter id para uuid: %v", err)
		}

		if userAuthorizedFromToken {
			userRow, err := conn.Collection(os.Getenv("MONGODB_COLLECTION_USER")).FindOne(ctx, bson.M{
				"id": userIdFromTokenUUID,
			}).DecodeBytes()

			if err != nil {
				return nil, fmt.Errorf("falha ao buscar usuário: %v", err)
			}

			_, data := userRow.Lookup("id").Binary()

			userDBId, err := uuid.FromBytes(data)

			if err != nil {
				return nil, fmt.Errorf("falha ao converter id para uuid: %v", err)
			}

			userDB, err := user.NewBuilder().WithID(userDBId).WithName(userRow.Lookup("name").StringValue()).WithPassword(userRow.Lookup("password").StringValue()).WithEmail(userRow.Lookup("email").StringValue()).WithAdmin(userRow.Lookup("admin").Boolean()).Build()

			if err != nil {
				return nil, err
			}

			var userType string

			if userDB.Admin() {
				userType = "admin"
				return &userType, nil
			} else {
				userType = "user"
				return &userType, nil
			}
		}

	} else {
		return nil, fmt.Errorf("token inválido")
	}

	userNotLogged := "not logged"

	return &userNotLogged, nil
}

func NewUserMongodbRepository(connectorManager connectorManager) *UserMongodbRepository {
	return &UserMongodbRepository{connectorManager: connectorManager}
}
