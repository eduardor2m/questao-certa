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

func TestCreateQuestion(t *testing.T) {
	controll := gomock.NewController(t)
	defer controll.Finish()

	repo := mocks.NewMockQuestionLoader(controll)
	service := NewQuestionServices(repo)

	id := uuid.New()
	baseReceived, err := base.NewBuilder().WithID(id).WithOrganization("Teste").WithDiscipline("Matemática").WithModel("multiple_choice").WithYear("2019").WithTopic("Qual o resultado de 1 + 1?").Build()
	assert.NoError(t, err, "Erro ao criar a base")

	questionReceived, err := question.NewBuilder().WithQuestion("Qual o resultado de 1 + 1?").WithAnswer("2").WithOptions([]string{"1", "2", "3", "4"}).Build()
	assert.NoError(t, err, "Erro ao criar a pergunta")
	questionReceived.Base = *baseReceived

	// Configurando expectativas no mock
	repo.EXPECT().CreateQuestion(questionReceived).Return(nil)

	err = service.CreateQuestion(*questionReceived)

	assert.NoError(t, err, "Erro ao criar a pergunta")
}
