package services

import (
	"encoding/csv"
	"errors"
	"mime/multipart"
	"strings"
	"sync"

	"github.com/eduardor2m/questao-certa/internal/app/entity/question"

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

func (instance *QuestionServices) Create(questionReceived question.Question) error {
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

	return instance.questionRepository.Create(*questionFormatted)
}

func (instance *QuestionServices) List(page int) ([]question.Question, error) {
	return instance.questionRepository.List(page)
}

func (instance *QuestionServices) ImportByCSV(file multipart.File) error {
	reader := csv.NewReader(file)
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	records = records[1:]

	var allQuestions []question.Question
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, record := range records {
		wg.Add(1)
		go func(record []string) {
			defer wg.Done()

			id, err := uuid.NewUUID()
			if err != nil {
				return
			}

			questionFormatted, err := question.NewBuilder().WithQuestion(record[5]).WithAnswer(record[6]).WithOptions(strings.Split(record[7], ";;")).Build()
			if err != nil {
				return
			}

			baseFormatted, err := base.NewBuilder().WithID(id).WithOrganization(record[0]).WithModel(record[1]).WithYear(record[2]).WithDiscipline(record[3]).WithTopic(record[4]).Build()
			if err != nil {
				return
			}

			questionFormatted.Base = *baseFormatted
			mu.Lock()
			allQuestions = append(allQuestions, *questionFormatted)
			mu.Unlock()
		}(record)
	}

	wg.Wait()

	return instance.questionRepository.ImportByCSV(allQuestions)
}

func (instance *QuestionServices) ListByFilter(filterReceived filter.Filter) ([]question.Question, error) {
	if filterReceived.Quantity() > 10 {
		return nil, errors.New("quantity must be less than 10")
	}

	return instance.questionRepository.ListByFilter(filterReceived)
}

func (instance *QuestionServices) DeleteByID(id string) error {
	return instance.questionRepository.DeleteByID(id)
}

func (instance *QuestionServices) DeleteAll() error {
	return instance.questionRepository.DeleteAll()
}

func NewQuestionServices(questionRepository repository.QuestionLoader) *QuestionServices {
	return &QuestionServices{questionRepository: questionRepository}
}
