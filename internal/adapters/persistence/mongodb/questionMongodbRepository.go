package mongodb

import (
	"context"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var _ repository.QuestionLoader = &QuestionMongodbRepository{}

type QuestionMongodbRepository struct {
	connectorManager
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
		"content":      Question.Content(),
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
		baseReceived, err := base.NewBuilder().WithID(question.ID).WithOrganization(question.Organization).WithModel(question.Model).WithYear(question.Year).WithContent(question.Content).WithTopic(question.Topic).Build()
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

func (instance *QuestionMongodbRepository) ListQuestionsByOrganization(organization string) ([]multiplechoice.MultipleChoice, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": organization})

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
		baseReceived, err := base.NewBuilder().WithID(question.ID).WithOrganization(question.Organization).WithModel(question.Model).WithYear(question.Year).WithContent(question.Content).WithTopic(question.Topic).Build()
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
