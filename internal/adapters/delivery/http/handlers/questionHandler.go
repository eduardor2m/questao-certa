package handlers

import (
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/response"
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
// @Param question body request.QuestionDTO true "Dados da questão de múltipla escolha"
// @Success 200 {object} response.InfoResponse "Questão criada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao criar questão"
// @Router /question [post]
func (instance QuestionHandler) CreateQuestion(context echo.Context) error {
	var questionDTO request.QuestionDTO

	err := context.Bind(&questionDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	baseReceived, err := base.NewBuilder().WithOrganization(questionDTO.Organization).WithModel(questionDTO.Model).WithYear(questionDTO.Year).WithDiscipline(questionDTO.Discipline).WithTopic(questionDTO.Topic).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	multipleChoiceReceived, err := question.NewBuilder().WithQuestion(questionDTO.Question).WithOptions(questionDTO.Options).WithAnswer(questionDTO.Answer).Build()
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

// @Summary Importa questões de múltipla escolha
// @Description Importa questões de múltipla escolha a partir de um arquivo CSV
// @Tags Question
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param file formData file true "Arquivo CSV com as questões de múltipla escolha"
// @Success 200 {object} response.InfoResponse
// @Failure 400 {object} response.Error
// @Router /question/import [post]
func (instance QuestionHandler) ImportQuestionsByCSV(context echo.Context) error {
	file, err := context.FormFile("file")
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	fileOpened, err := file.Open()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.ImportQuestionsByCSV(fileOpened)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Questions imported successfully"})
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
	page := context.Param("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}
	questionsReceivedDB, err := instance.service.ListQuestions(pageInt)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var questionsDTO []response.Question

	for _, questionReceivedDB := range questionsReceivedDB {
		questionDTO := response.Question{
			ID:           questionReceivedDB.ID(),
			Organization: questionReceivedDB.Organization(),
			Model:        questionReceivedDB.Model(),
			Year:         questionReceivedDB.Year(),
			Discipline:   questionReceivedDB.Discipline(),
			Topic:        questionReceivedDB.Topic(),
			Question:     questionReceivedDB.Question(),
			Options:      questionReceivedDB.Options(),
			Answer:       questionReceivedDB.Answer(),
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

	filterFormatted, err := filter.NewBuilder().WithOrganization(filterReceived.Organization).WithYear(filterReceived.Year).WithDiscipline(filterReceived.Discipline).WithTopic(filterReceived.Topic).WithQuantity(filterReceived.Quantity).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	questionsReceivedDB, err := instance.service.ListQuestionsByFilter(*filterFormatted)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var questionsDTO []response.Question

	for _, questionReceivedDB := range questionsReceivedDB {
		questionDTO := response.Question{
			ID:           questionReceivedDB.ID(),
			Organization: questionReceivedDB.Organization(),
			Model:        questionReceivedDB.Model(),
			Year:         questionReceivedDB.Year(),
			Discipline:   questionReceivedDB.Discipline(),
			Topic:        questionReceivedDB.Topic(),
			Question:     questionReceivedDB.Question(),
			Options:      questionReceivedDB.Options(),
			Answer:       questionReceivedDB.Answer(),
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
