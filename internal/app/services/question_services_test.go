package services

import (
	"testing"

	"github.com/eduardor2m/questao-certa/internal/app/entity/filter"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question"
	"github.com/eduardor2m/questao-certa/internal/app/entity/question/base"
	"github.com/eduardor2m/questao-certa/internal/app/services/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateQuestion(t *testing.T) {
	t.Skip("TestCreateQuestion not implemented")
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// repo := mocks.NewMockQuestionLoader(ctrl)
	// questionService := NewQuestionServices(repo)

	// questionMock, err := question.NewBuilder().WithQuestion("Question").WithAnswer("(A) Option 1").WithOptions([]string{"(A) Option 1", "(B) Option 2", "(C) Option 3", "(D) Option 4", "(E) Option 5"}).Build()
	// if err != nil {
	// 	t.Error(err)
	// }

	// baseMock, err := base.NewBuilder().WithID(uuid.New()).WithOrganization("Organization").WithModel("true_or_false").WithYear("2020").WithDiscipline("Discipline").WithTopic("Topic").Build()
	// if err != nil {
	// 	t.Error(err)
	// }

	// questionMock.Base = *baseMock

	// repo.EXPECT().CreateQuestion(gomock.Any()).Return(assert.AnError)

	// err = questionService.CreateQuestion(*questionMock)

	// assert.Nil(t, err)
}

func TestListQuestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	page := 1

	repo := mocks.NewMockQuestionLoader(ctrl)
	questionService := NewQuestionServices(repo)

	var questionsMock []question.Question

	for i := 0; i < 3; i++ {
		questionMock, err := question.NewBuilder().WithQuestion("Question").WithAnswer("(A) Option 1").WithOptions([]string{"(A) Option 1", "(B) Option 2", "(C) Option 3", "(D) Option 4", "(E) Option 5"}).Build()
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
	assert.Equal(t, questions[0].Answer(), "(A) Option 1")
	assert.Equal(t, questions[0].Options(), []string{"(A) Option 1", "(B) Option 2", "(C) Option 3", "(D) Option 4", "(E) Option 5"})
}

func TestListQuestionsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	page := 1

	repo := mocks.NewMockQuestionLoader(ctrl)
	questionService := NewQuestionServices(repo)

	repo.EXPECT().ListQuestions(page).Return(nil, assert.AnError)

	questions, err := questionService.ListQuestions(page)

	assert.NotNil(t, err)
	assert.Nil(t, questions)
}

func TestListQuestionsByFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockQuestionLoader(ctrl)
	questionService := NewQuestionServices(repo)

	var questionsMock []question.Question

	for i := 0; i < 3; i++ {
		questionMock, err := question.NewBuilder().WithQuestion("Question").WithAnswer("(A) Option 1").WithOptions([]string{"(A) Option 1", "(B) Option 2", "(C) Option 3", "(D) Option 4", "(E) Option 5"}).Build()
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

	f, err := filter.NewBuilder().WithOrganization("Organization").WithYear("2020").WithDiscipline("Discipline").WithTopic("Topic").WithQuantity(1).Build()
	if err != nil {
		t.Error(err)
	}

	repo.EXPECT().ListQuestionsByFilter(*f).Return(questionsMock, nil)

	questions, err := questionService.ListQuestionsByFilter(*f)

	assert.Nil(t, err)
	assert.NotNil(t, questions)
	assert.Equal(t, len(questions), 3)
	assert.Equal(t, questions[0].Question(), "Question")
	assert.Equal(t, questions[0].Answer(), "(A) Option 1")
	assert.Equal(t, questions[0].Options(), []string{"(A) Option 1", "(B) Option 2", "(C) Option 3", "(D) Option 4", "(E) Option 5"})
}
