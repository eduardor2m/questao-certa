package handlers

import (
	"net/http"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/response"
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
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

	baseReceived, err := base.NewBuilder().WithOrganization(multipleChoiceDTO.Organization).WithModel(multipleChoiceDTO.Model).WithYear(multipleChoiceDTO.Year).WithContent(multipleChoiceDTO.Content).WithTopic(multipleChoiceDTO.Topic).Build()

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

// @Summary Lista todas as questões
// @Description Lista todas as questões
// @Tags Question
// @Accept json
// @Produce json
// @Security bearerAuth
// @Success 200 {array} response.MultipleChoice
// @Failure 400 {object} response.Error
// @Router /question [get]
func (instance QuestionHandler) ListQuestions(context echo.Context) error {
	questions, err := instance.service.ListQuestions()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var questionsDTO []response.MultipleChoice

	for _, question := range questions {
		questionDTO := response.MultipleChoice{
			ID:           question.ID(),
			Organization: question.Organization(),
			Model:        question.Model(),
			Year:         question.Year(),
			Content:      question.Content(),
			Topic:        question.Topic(),
			Question:     question.Question(),
			Options:      question.Options(),
			Answer:       question.Answer(),
		}

		questionsDTO = append(questionsDTO, questionDTO)
	}

	return context.JSON(http.StatusOK, questionsDTO)
}

// @Summary Lista todas as questões de uma organização
// @Description Lista todas as questões de uma organização
// @Tags Question
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param organization path string true "Nome da organização"
// @Success 200 {array} response.MultipleChoice
// @Failure 400 {object} response.Error
// @Router /question/{organization} [get]
func (instance QuestionHandler) ListQuestionsByFilter(context echo.Context) error {
	filterReceived := request.FilterDTO{}

	err := context.Bind(&filterReceived)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	filterFormated, err := filter.NewBuilder().WithOrganization(filterReceived.Organization).WithYear(filterReceived.Year).WithContent(filterReceived.Content).WithTopic(filterReceived.Topic).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	questions, err := instance.service.ListQuestionsByFilter(*filterFormated)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var questionsDTO []response.MultipleChoice

	for _, question := range questions {
		questionDTO := response.MultipleChoice{
			ID:           question.ID(),
			Organization: question.Organization(),
			Model:        question.Model(),
			Year:         question.Year(),
			Content:      question.Content(),
			Topic:        question.Topic(),
			Question:     question.Question(),
			Options:      question.Options(),
			Answer:       question.Answer(),
		}

		questionsDTO = append(questionsDTO, questionDTO)
	}

	return context.JSON(http.StatusOK, questionsDTO)
}

// @Summary Deleta uma questão
// @Description Deleta uma questão
// @Tags Question
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param id path string true "ID da questão"
// @Success 200 {object} response.Info
// @Failure 400 {object} response.Error
// @Router /question/{id} [delete]
func (instance QuestionHandler) DeleteQuestion(context echo.Context) error {
	id := context.Param("id")

	err := instance.service.DeleteQuestion(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Question deleted successfully"})
}

// @Summary Deleta todas as questões
// @Description Deleta todas as questões
// @Tags Question
// @Accept json
// @Produce json
// @Security bearerAuth
// @Success 200 {object} response.Info
// @Failure 400 {object} response.Error
// @Router /question [delete]
func (instance QuestionHandler) DeleteAllQuestions(context echo.Context) error {
	err := instance.service.DeleteAllQuestions()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Questions deleted successfully"})
}

func NewQuestionHandler(service primary.QuestionManager) *QuestionHandler {
	return &QuestionHandler{
		service: service,
	}
}
