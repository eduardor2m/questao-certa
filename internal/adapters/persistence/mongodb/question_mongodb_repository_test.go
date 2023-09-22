package mongodb_test

// import (
// 	"os"
// 	"testing"

// 	"github.com/eduardor2m/questao-certa/internal/adapters/persistence/mongodb"
// 	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
// )

// func TestListQuestions(t *testing.T) {
// 	dbName := "database_test"
// 	collectionName := "questions_test"

// 	err := os.Setenv("DB_COLLECTION", collectionName)
// 	if err != nil {
// 		t.Errorf("error setting env variable: %v", err)
// 	}
// 	defer os.Unsetenv("DB_COLLECTION")

// 	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer mtestDB.Close()

// 	mtestDB.Run("TestListQuestions", func(mt *mtest.T) {
// 		mt.AddMockResponses(mtest.CreateCursorResponse(0, "questions_test", "questions_test", mtest.FirstBatch, mtest.NoError))
// 		mt.AddMockResponses(mtest.CreateSuccessResponse())

// 		databaseMock := mt.Client.Database(dbName)

// 		repo := mongodb.NewQuestionMongodbRepository(databaseMock)

// 		_, err := repo.ListQuestions(1)
// 		if err != nil {
// 			t.Errorf("error listing questions: %v", err)
// 		}
// 	})

// }
