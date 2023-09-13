package services

import (
	"testing"

	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/tools/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestListQuestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	page := 1

	repo := mocks.NewMockQuestionLoader(ctrl)
	questionService := NewQuestionServices(repo)

	var questionsMock []question.Question

	for i := 0; i < 3; i++ {
		questionMock, err := question.NewBuilder().WithQuestion("Question").WithAnswer("Answer").WithOptions([]string{"Option 1", "Option 2", "Option 3"}).Build()
		if err != nil {
			t.Error(err)
		}

		baseMock, err := base.NewBuilder().WithID(uuid.New()).WithOrganization("Organization").WithModel("true_or_false").WithYear("2020").WithDiscipline("Discipline").WithTopic("Topic").Build()
		if err != nil {
			t.Error(err)
		}

		questionMock.Base = *baseMock

		questionsMock = append(questionsMock, *questionMock)
	}

	repo.EXPECT().ListQuestions(page).Return(questionsMock, nil)

	questions, err := questionService.ListQuestions(page)

	assert.Nil(t, err)
	assert.NotNil(t, questions)
	assert.Equal(t, len(questions), 3)
	assert.Equal(t, questions[0].Question(), "Question")
	assert.Equal(t, questions[0].Answer(), "Answer")
	assert.Equal(t, questions[0].Options(), []string{"Option 1", "Option 2", "Option 3"})
}
