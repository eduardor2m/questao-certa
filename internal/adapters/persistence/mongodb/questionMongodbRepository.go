package mongodb

import (
	"context"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/handlers/dto/request"
	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	multiplechoice "github.com/eduardor2m/questao-certa/internal/app/entity/question/multipleChoice"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (instance *QuestionMongodbRepository) ListQuestionsByFilter(f filter.Filter) ([]multiplechoice.MultipleChoice, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	if f.Organization() == "" && f.Year() == "" && f.Topic() == "" && f.Content() == "" {
		return []multiplechoice.MultipleChoice{}, nil
	} else if f.Organization() == "" && f.Year() == "" && f.Topic() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Organization() == "" && f.Year() == "" && f.Content() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"topic": f.Topic()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Organization() == "" && f.Topic() == "" && f.Content() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"year": f.Year()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Year() == "" && f.Topic() == "" && f.Content() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Organization() == "" && f.Year() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"topic": f.Topic(), "content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Organization() == "" && f.Topic() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"year": f.Year(), "content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Organization() == "" && f.Content() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"year": f.Year(), "topic": f.Topic()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Year() == "" && f.Topic() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization(), "content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Year() == "" && f.Content() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization(), "topic": f.Topic()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Topic() == "" && f.Content() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization(), "year": f.Year()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Organization() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"year": f.Year(), "topic": f.Topic(), "content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Year() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization(), "topic": f.Topic(), "content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Topic() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization(), "year": f.Year(), "content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	} else if f.Content() == "" {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization(), "year": f.Year(), "topic": f.Topic()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)
	} else {
		cursor, err := conn.Collection("questions").Find(ctx, bson.M{"organization": f.Organization(), "year": f.Year(), "topic": f.Topic(), "content": f.Content()})
		if err != nil {
			return []multiplechoice.MultipleChoice{}, err
		}

		return formater(cursor, ctx)

	}

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
