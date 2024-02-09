package main_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/eduardor2m/questao-certa/tools/tests/data"
	"github.com/eduardor2m/questao-certa/tools/tests/data/jsons/filenames"
	"github.com/eduardor2m/questao-certa/tools/tests/e2e/utils"
)

const baseURL = "http://localhost:8180"

func TestCreateQuestion(t *testing.T) {
	questionData := data.GetQuestionMock(filenames.QuestionMock)
	requestBody := bytes.NewReader(questionData)

	clientRequest, err := http.NewRequest("POST", "http://questao-certa-air:8080/api/question", requestBody)

	if err != nil {
		t.Errorf("Erro ao criar a solicitação: %v", err)
		return
	}

	token, err := utils.GenerateToken()
	if err != nil {
		t.Errorf("Erro ao gerar token: %v", err)
		return
	}

	clientRequest.Header.Set("Content-Type", "application/json")
	clientRequest.Header.Set("Authorization", "Bearer "+*token)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Errorf("Erro ao fazer a solicitação: %v", err)
		return
	}
	defer serverResponse.Body.Close()

	if serverResponse.StatusCode != http.StatusCreated {
		t.Errorf("Esperado código de status %d, mas obteve %d", http.StatusCreated, serverResponse.StatusCode)
		return
	}
}

func TestGetQuestions(t *testing.T) {
	clientRequest, err := http.NewRequest("GET", "http://questao-certa-air:8080/api/question/1", nil)

	if err != nil {
		t.Errorf("Erro ao criar a solicitação: %v", err)
		return
	}

	token, err := utils.GenerateToken()

	if err != nil {
		t.Errorf("Erro ao gerar token: %v", err)
		return
	}

	clientRequest.Header.Set("Content-Type", "application/json")

	clientRequest.Header.Set("Authorization", "Bearer "+*token)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Errorf("Erro ao fazer a solicitação: %v", err)
		return
	}

	defer serverResponse.Body.Close()

	if serverResponse.StatusCode != http.StatusOK {
		t.Errorf("Esperado código de status %d, mas obteve %d", http.StatusOK, serverResponse.StatusCode)
		return
	}
}

func TestFilterQuestions(t *testing.T) {
	filterData := data.GetFilterMock(filenames.FilterMock)
	requestBody := bytes.NewReader(filterData)
	clientRequest, err := http.NewRequest("POST", "http://questao-certa-air:8080/api/question/filter", requestBody)

	if err != nil {
		t.Errorf("Erro ao criar a solicitação: %v", err)
		return
	}

	token, err := utils.GenerateToken()

	if err != nil {
		t.Errorf("Erro ao gerar token: %v", err)
		return
	}

	clientRequest.Header.Set("Content-Type", "application/json")

	clientRequest.Header.Set("Authorization", "Bearer "+*token)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Errorf("Erro ao fazer a solicitação: %v", err)
		return
	}

	defer serverResponse.Body.Close()

	if serverResponse.StatusCode != http.StatusOK {
		t.Errorf("Esperado código de status %d, mas obteve %d", http.StatusOK, serverResponse.StatusCode)
		return
	}
}

func TestUpdateQuestion(t *testing.T) {
	t.Skip("Teste não implementado")
}

func TestDeleteQuestion(t *testing.T) {
	t.Skip("Teste não implementado")
}
