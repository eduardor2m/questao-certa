package mongodb

import (
	"context"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ repository.QuestionLoader = &QuestionMongodbRepository{}

type QuestionMongodbRepository struct {
	connectorManager
}

func formater(c *mongo.Cursor, ctx context.Context) ([]multiplechoice.MultipleChoice, error) {
	var questions []request.MultipleChoiceDTO

	err := c.All(ctx, &questions)
	if err != nil {
		return nil, err
	}

	var multipleChoicesDB []multiplechoice.MultipleChoice

	for _, question := range questions {
		baseReceived, err := base.NewBuilder().WithID(question.ID).WithOrganization(question.Organization).WithModel(question.Model).WithYear(question.Year).WithDiscipline(question.Discipline).WithTopic(question.Topic).Build()

		if err != nil {
			return nil, err
		}

		multipleChoice, err := multiplechoice.NewBuilder().WithQuestion(question.Question).WithAnswer(question.Answer).WithOptions(question.Options).Build()
		if err != nil {
			return nil, err
		}

		multipleChoice.Base = *baseReceived

		multipleChoicesDB = append(multipleChoicesDB, *multipleChoice)
	}

	return multipleChoicesDB, nil
}

func (instance *QuestionMongodbRepository) CreateQuestion(Question multiplechoice.MultipleChoice) error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	ctx := context.Background()

	document := bson.M{
		"id":           Question.ID(),
		"organization": Question.Organization(),
		"model":        Question.Model(),
		"year":         Question.Year(),
		"discipline":   Question.Discipline(),
		"topic":        Question.Topic(),
		"question":     Question.Question(),
		"answer":       Question.Answer(),
		"options":      Question.Options(),
	}

	_, err = conn.Collection("questions").InsertOne(ctx, document)
	if err != nil {
		return err
	}

	return nil
}

func (instance *QuestionMongodbRepository) ImportQuestionsByCSV(questions []multiplechoice.MultipleChoice) error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	ctx := context.Background()

	var documents []interface{}

	for _, question := range questions {
		document := bson.M{
			"id":           question.ID(),
			"organization": question.Organization(),
			"model":        question.Model(),
			"year":         question.Year(),
			"discipline":   question.Discipline(),
			"topic":        question.Topic(),
			"question":     question.Question(),
			"answer":       question.Answer(),
			"options":      question.Options(),
		}

		documents = append(documents, document)
	}

	_, err = conn.Collection("questions").InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	return nil
}

func (instance *QuestionMongodbRepository) ListQuestions() ([]multiplechoice.MultipleChoice, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	cursor, err := conn.Collection("questions").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var questions []request.MultipleChoiceDTO
	err = cursor.All(ctx, &questions)
	if err != nil {
		return nil, err
	}

	var multipleChoicesDB []multiplechoice.MultipleChoice

	for _, question := range questions {
		baseReceived, err := base.NewBuilder().WithID(question.ID).WithOrganization(question.Organization).WithModel(question.Model).WithYear(question.Year).WithDiscipline(question.Discipline).WithTopic(question.Topic).Build()
		if err != nil {
			return nil, err
		}

		multipleChoice, err := multiplechoice.NewBuilder().WithQuestion(question.Question).WithAnswer(question.Answer).WithOptions(question.Options).Build()
		if err != nil {
			return nil, err
		}

		multipleChoice.Base = *baseReceived

		multipleChoicesDB = append(multipleChoicesDB, *multipleChoice)
	}

	return multipleChoicesDB, nil
}

func (instance *QuestionMongodbRepository) ListQuestionsByFilter(f filter.Filter) ([]multiplechoice.MultipleChoice, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	findOptions := options.Find()

	filterQuery := bson.M{}

	if f.Quantity() != 0 {
		findOptions.SetLimit(f.Quantity())
	} else {
		findOptions.SetLimit(3)
	}

	if f.Organization() != "" {
		filterQuery["organization"] = f.Organization()
	}
	if f.Year() != "" {
		filterQuery["year"] = f.Year()
	}
	if f.Topic() != "" {
		filterQuery["topic"] = f.Topic()
	}
	if f.Discipline() != "" {
		filterQuery["discipline"] = f.Discipline()
	}

	cursor, err := conn.Collection("questions").Find(ctx, filterQuery, findOptions)
	if err != nil {
		return []multiplechoice.MultipleChoice{}, err
	}

	return formater(cursor, ctx)
}

func (instance *QuestionMongodbRepository) DeleteQuestion(id string) error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = conn.Collection("questions").DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}

func (instance *QuestionMongodbRepository) DeleteAllQuestions() error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = conn.Collection("questions").DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func NewQuestionMongodbRepository(connectorManager connectorManager) *QuestionMongodbRepository {
	return &QuestionMongodbRepository{connectorManager: connectorManager}
}
