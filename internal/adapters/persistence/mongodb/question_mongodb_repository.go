package mongodb

import (
	"context"
	"os"

	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/internal/app/interfaces/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ repository.QuestionLoader = &QuestionMongodbRepository{}

type QuestionMongodbRepository struct {
	connectorManager
}

type QuestionDB struct {
	ID           uuid.UUID `bson:"id"`
	Organization string    `bson:"organization"`
	Model        string    `bson:"model"`
	Year         string    `bson:"year"`
	Discipline   string    `bson:"discipline"`
	Topic        string    `bson:"topic"`
	Question     string    `bson:"question"`
	Answer       string    `bson:"answer"`
	Options      []string  `bson:"options"`
}

func mapBSONToQuestions(c *mongo.Cursor, ctx context.Context) ([]question.Question, error) {
	var questionsDB []QuestionDB

	err := c.All(ctx, &questionsDB)
	if err != nil {
		return nil, err
	}

	var questionFormattedDB []question.Question

	for _, questionDB := range questionsDB {
		baseFormatted, err := base.NewBuilder().WithID(questionDB.ID).WithOrganization(questionDB.Organization).WithModel(questionDB.Model).WithYear(questionDB.Year).WithDiscipline(questionDB.Discipline).WithTopic(questionDB.Topic).Build()

		if err != nil {
			return nil, err
		}

		questionFormatted, err := question.NewBuilder().WithQuestion(questionDB.Question).WithAnswer(questionDB.Answer).WithOptions(questionDB.Options).Build()
		if err != nil {
			return nil, err
		}

		questionFormatted.Base = *baseFormatted

		questionFormattedDB = append(questionFormattedDB, *questionFormatted)
	}

	return questionFormattedDB, nil
}

func (instance *QuestionMongodbRepository) CreateQuestion(questionReceived question.Question) error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	document := bson.M{
		"id":           questionReceived.ID(),
		"organization": questionReceived.Organization(),
		"model":        questionReceived.Model(),
		"year":         questionReceived.Year(),
		"discipline":   questionReceived.Discipline(),
		"topic":        questionReceived.Topic(),
		"question":     questionReceived.Question(),
		"answer":       questionReceived.Answer(),
		"options":      questionReceived.Options(),
	}

	collectionName := os.Getenv("MONGODB_COLLECTION")

	_, err = conn.Collection(collectionName).InsertOne(ctx, document)
	if err != nil {
		return err
	}

	return nil
}

func (instance *QuestionMongodbRepository) ImportQuestionsByCSV(questionsReceived []question.Question) error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	ctx := context.Background()

	var documents []interface{}

	for _, questionReceived := range questionsReceived {
		document := bson.M{
			"id":           questionReceived.ID(),
			"organization": questionReceived.Organization(),
			"model":        questionReceived.Model(),
			"year":         questionReceived.Year(),
			"discipline":   questionReceived.Discipline(),
			"topic":        questionReceived.Topic(),
			"question":     questionReceived.Question(),
			"answer":       questionReceived.Answer(),
			"options":      questionReceived.Options(),
		}

		documents = append(documents, document)
	}
	collectionName := os.Getenv("MONGODB_COLLECTION")

	_, err = conn.Collection(collectionName).InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	return nil
}

func (instance *QuestionMongodbRepository) ListQuestions(page int) ([]question.Question, error) {
	ctx := context.Background()
	perPage := 3
	collectionName := os.Getenv("MONGODB_COLLECTION")

	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}
	defer instance.closeConnection(conn)

	findOptions := options.Find().
		SetLimit(int64(perPage)).
		SetSkip(int64(perPage * (page - 1)))

	cursor, err := conn.Collection(collectionName).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	return mapBSONToQuestions(cursor, ctx)
}

func (instance *QuestionMongodbRepository) ListQuestionsByFilter(f filter.Filter) ([]question.Question, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}

	defer instance.closeConnection(conn)

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

	collectionName := os.Getenv("MONGODB_COLLECTION")

	cursor, err := conn.Collection(collectionName).Find(ctx, filterQuery, findOptions)
	if err != nil {
		return nil, err
	}

	return mapBSONToQuestions(cursor, ctx)
}

func (instance *QuestionMongodbRepository) DeleteQuestion(id string) error {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	collectionName := os.Getenv("MONGODB_COLLECTION")

	_, err = conn.Collection(collectionName).DeleteOne(ctx, bson.M{"id": id})
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

	defer instance.closeConnection(conn)

	ctx := context.Background()

	collectionName := os.Getenv("MONGODB_COLLECTION")

	_, err = conn.Collection(collectionName).DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func NewQuestionMongodbRepository(connectorManager connectorManager) *QuestionMongodbRepository {
	return &QuestionMongodbRepository{connectorManager: connectorManager}
}
