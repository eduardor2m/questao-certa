package handlers

import (
	"net/http"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/response"
	"github.com/eduardor2m/questao-certa/internal/app/entity/user"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service primary.UserManager
}

// @Summary Cria um usuário
// @Description Cria um usuário
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserDTO true "Dados do usuário"
// @Success 201 {object} response.InfoResponse "Usuário criado com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao criar usuário"
// @Router /user [post]
func (instance UserHandler) Register(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	userReceiced, err := user.NewBuilder().WithName(userDTO.Name).WithPassword(userDTO.Password).WithEmail(userDTO.Email).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.Register(*userReceiced)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User created successfully"})
}

// @Summary Autentica um usuário
// @Description Autentica um usuário
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserLoginDTO true "Dados do usuário"
// @Success 200 {object} response.InfoResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /user/signin [post]
func (instance UserHandler) Authenticate(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	token, err := instance.service.Authenticate(userDTO.Email, userDTO.Password)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	context.Response().Header().Set("Authorization", *token)
	context.Response().Header().Set("Access-Control-Expose-Headers", "Authorization")

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User logged successfully"})
}

// @Summary Deleta um usuário
// @Description Deleta um usuário
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserDeleteDTO true "Dados do usuário"
// @Success 200 {object} response.InfoResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /user [delete]
func (instance UserHandler) Delete(context echo.Context) error {
	var userDTO request.UserDeleteDTO

	err := context.Bind(&userDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.Delete(userDTO.Email, userDTO.Name)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User deleted successfully"})
}

// @Summary Lista os usuários
// @Description Lista os usuários
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} []response.UserDTO
// @Failure 400 {object} response.ErrorResponse
// @Router /user [get]
func (instance UserHandler) List(context echo.Context) error {
	users, err := instance.service.List()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var usersJson []response.UserDTO

	for _, user := range users {
		userJson := response.UserDTO{
			ID:        user.ID(),
			Name:      user.Name(),
			Email:     user.Email(),
			Password:  user.Password(),
			Admin:     user.Admin(),
			IsActive:  user.IsActive(),
			CreatedAt: user.CreatedAt().Format("2006-01-02T15:04:05Z"),
			UpdatedAt: user.UpdatedAt().Format("2006-01-02T15:04:05Z"),
		}

		usersJson = append(usersJson, userJson)
	}

	return context.JSON(http.StatusOK, usersJson)
}

// @Summary Busca um usuário pelo email
// @Description Busca um usuário pelo email
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "Email do usuário"
// @Success 200 {object} response.UserDTO
// @Failure 400 {object} response.ErrorResponse
// @Router /user/{email} [get]
func (instance UserHandler) FindByEmail(context echo.Context) error {
	email := context.Param("email")

	user, err := instance.service.FindByEmail(email)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	userJson := response.UserDTO{
		ID:        user.ID(),
		Name:      user.Name(),
		Email:     user.Email(),
		Password:  user.Password(),
		Admin:     user.Admin(),
		CreatedAt: user.CreatedAt().Format("2006-01-02T15:04:05Z"),
		UpdatedAt: user.UpdatedAt().Format("2006-01-02T15:04:05Z"),
	}

	return context.JSON(http.StatusOK, userJson)
}

// @Summary Verifica o tipo do usuário
// @Description Verifica o tipo do usuário
// @Tags User
// @Accept json
// @Produce json
// @Param token header string true "Token de autenticação"
// @Success 200 {object} response.InfoResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /user/verify [get]
func (instance UserHandler) CheckType(context echo.Context) error {
	token := context.Request().Header.Get("Authorization")

	if token == "" {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Token not found"})
	}

	tokenFormatted := token[7:]

	message, err := instance.service.CheckType(tokenFormatted)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "user type is: " + *message})
}

func NewUserHandler(service primary.UserManager) *UserHandler {
	return &UserHandler{service: service}
}
