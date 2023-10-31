package main_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/eduardor2m/questao-certa/tools/tests/data"
	"github.com/eduardor2m/questao-certa/tools/tests/data/jsons/filenames"
)

func TestCreateQuestion(t *testing.T) {
	questionData := data.GetQuestionMock(filenames.QuestionMock)
	requestBody := bytes.NewReader(questionData)

	clientRequest, err := http.NewRequest("POST", "http://localhost:8180/api/question", requestBody)

	if err != nil {
		t.Errorf("Erro ao criar a solicitação: %v", err)
		return
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTg3Mzc4ODAsInVzZXJfaWQiOiIyYzViNTdjZS03NTM0LTExZWUtYjAxNC0wMjQyYWMxMjAwMDUifQ.VkPWCPETYHDUFAY_AcqSWvzkseDb-GmMy_6Xm_kJyqs"

	clientRequest.Header.Set("Content-Type", "application/json")
	clientRequest.Header.Set("Authorization", "Bearer "+token)

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

func TestGetQuestion(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetQuestions(t *testing.T) {
	t.Skip("Not implemented")
}

func TestUpdateQuestion(t *testing.T) {
	t.Skip("Not implemented")
}

func TestDeleteQuestion(t *testing.T) {
	t.Skip("Not implemented")
}
