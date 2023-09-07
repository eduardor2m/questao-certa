package handlers

import (
	"net/http"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/response"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/labstack/echo/v4"
)

type QuestionHandler struct {
	service primary.QuestionManager
}

// @Summary Cria uma questão de múltipla escolha
// @Description Cria uma questão de múltipla escolha, onde só é possível uma resposta correta
// @Tags Question
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param question body request.MultipleChoiceDTO true "Dados da questão de múltipla escolha"
// @Success 200 {object} response.InfoResponse "Questão criada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao criar questão"
// @Router /question [post]
func (instance QuestionHandler) CreateQuestion(context echo.Context) error {
	var multipleChoiceDTO request.MultipleChoiceDTO

	err := context.Bind(&multipleChoiceDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	baseReceived, err := base.NewBuilder().WithID(multipleChoiceDTO.ID).WithOrganization(multipleChoiceDTO.Organization).WithModel(multipleChoiceDTO.Model).WithYear(multipleChoiceDTO.Year).WithContent(multipleChoiceDTO.Content).WithTopic(multipleChoiceDTO.Topic).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	multipleChoiceReceived, err := multiplechoice.NewBuilder().WithQuestion(multipleChoiceDTO.Question).WithOptions(multipleChoiceDTO.Options).WithAnswer(multipleChoiceDTO.Answer).Build()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	multipleChoiceReceived.Base = *baseReceived

	err = instance.service.CreateQuestion(*multipleChoiceReceived)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Question created successfully"})
}

func NewQuestionHandler(service primary.QuestionManager) *QuestionHandler {
	return &QuestionHandler{
		service: service,
	}
}
