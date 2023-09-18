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
func (instance UserHandler) SignUp(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	userReceiced, err := user.NewBuilder().WithName(userDTO.Name).WithPassword(userDTO.Password).WithEmail(userDTO.Email).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.SignUp(*userReceiced)

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
// @Param user body request.UserDTO true "Dados do usuário"
// @Success 200 {object} response.InfoResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /user/signin [post]
func (instance UserHandler) SignIn(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	token, err := instance.service.SignIn(userDTO.Email, userDTO.Password)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	context.Response().Header().Set("Authorization", *token)

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User logged successfully"})
}

func NewUserHandler(service primary.UserManager) *UserHandler {
	return &UserHandler{service: service}
}
