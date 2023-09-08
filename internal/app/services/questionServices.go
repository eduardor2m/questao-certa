package services

import (
	"encoding/csv"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/primary"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"github.com/google/uuid"
)

var _ primary.QuestionManager = (*QuestionServices)(nil)

type QuestionServices struct {
	questionRepository repository.QuestionLoader
}

func (instance *QuestionServices) CreateQuestion(question multiplechoice.MultipleChoice) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	questionFormated, err := multiplechoice.NewBuilder().WithQuestion(question.Question()).WithAnswer(question.Answer()).WithOptions(question.Options()).Build()
	if err != nil {
		return err
	}

	baseFormated, err := base.NewBuilder().WithID(id).WithOrganization(question.Organization()).WithModel(question.Model()).WithYear(question.Year()).WithContent(question.Content()).WithTopic(question.Topic()).Build()
	if err != nil {
		return err
	}

	questionFormated.Base = *baseFormated
	return instance.questionRepository.CreateQuestion(*questionFormated)
}

func (instance *QuestionServices) ListQuestions() ([]multiplechoice.MultipleChoice, error) {
	return instance.questionRepository.ListQuestions()
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

		questionFormated, err := multiplechoice.NewBuilder().WithQuestion(record[0]).WithAnswer(record[1]).WithOptions(strings.Split(record[2], ",")).Build()
		if err != nil {
			return err
		}

		baseFormated, err := base.NewBuilder().WithID(id).WithOrganization(record[3]).WithModel(record[4]).WithYear(record[5]).WithContent(record[6]).WithTopic(record[7]).Build()
		if err != nil {
			return err
		}

		questionFormated.Base = *baseFormated

		err = instance.questionRepository.CreateQuestion(*questionFormated)

		if err != nil {
			return err
		}

	}

	return nil
}

func (instance *QuestionServices) ListQuestionsByFilter(f filter.Filter) ([]multiplechoice.MultipleChoice, error) {
	if f.Quantity() > 10 {
		return nil, errors.New("quantity must be less than 10")
	}

	return instance.questionRepository.ListQuestionsByFilter(f)
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
