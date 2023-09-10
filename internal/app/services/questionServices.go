package services

import (
	"encoding/csv"
	"errors"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"mime/multipart"
	"strings"

	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"github.com/google/uuid"
)

var _ primary.QuestionManager = (*QuestionServices)(nil)

type QuestionServices struct {
	questionRepository repository.QuestionLoader
}

func (instance *QuestionServices) CreateQuestion(questionReceived question.Question) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	questionFormatted, err := question.NewBuilder().WithQuestion(questionReceived.Question()).WithAnswer(questionReceived.Answer()).WithOptions(questionReceived.Options()).Build()
	if err != nil {
		return err
	}

	baseFormatted, err := base.NewBuilder().WithID(id).WithOrganization(questionReceived.Organization()).WithModel(questionReceived.Model()).WithYear(questionReceived.Year()).WithDiscipline(questionReceived.Discipline()).WithTopic(questionReceived.Topic()).Build()
	if err != nil {
		return err
	}

	questionFormatted.Base = *baseFormatted

	return instance.questionRepository.CreateQuestion(*questionFormatted)
}

func (instance *QuestionServices) ListQuestions(page int) ([]question.Question, error) {
	return instance.questionRepository.ListQuestions(page)
}

func (instance *QuestionServices) ImportQuestionsByCSV(file multipart.File) error {
	reader := csv.NewReader(file)
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	records = records[1:]

	for _, record := range records {

		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}

		questionFormatted, err := question.NewBuilder().WithQuestion(record[5]).WithAnswer(record[6]).WithOptions(strings.Split(record[7], ";;")).Build()
		if err != nil {
			return err
		}

		baseFormatted, err := base.NewBuilder().WithID(id).WithOrganization(record[0]).WithModel(record[1]).WithYear(record[2]).WithDiscipline(record[3]).WithTopic(record[4]).Build()
		if err != nil {
			return err
		}

		questionFormatted.Base = *baseFormatted

		err = instance.questionRepository.CreateQuestion(*questionFormatted)

		if err != nil {
			return err
		}

	}

	return nil
}

func (instance *QuestionServices) ListQuestionsByFilter(filterReceived filter.Filter) ([]question.Question, error) {
	if filterReceived.Quantity() > 10 {
		return nil, errors.New("quantity must be less than 10")
	}

	return instance.questionRepository.ListQuestionsByFilter(filterReceived)
}

func (instance *QuestionServices) DeleteQuestion(id string) error {
	return instance.questionRepository.DeleteQuestion(id)
}

func (instance *QuestionServices) DeleteAllQuestions() error {
	return instance.questionRepository.DeleteAllQuestions()
}

func NewQuestionServices(questionRepository repository.QuestionLoader) *QuestionServices {
	return &QuestionServices{questionRepository: questionRepository}
}
