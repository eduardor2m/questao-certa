package mongodb

import (
	"context"

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

func NewQuestionMongodbRepository(connectorManager connectorManager) *QuestionMongodbRepository {
	return &QuestionMongodbRepository{connectorManager: connectorManager}
}
