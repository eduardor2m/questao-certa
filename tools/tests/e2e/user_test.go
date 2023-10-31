package main_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/eduardor2m/questao-certa/tools/tests/data"
	"github.com/eduardor2m/questao-certa/tools/tests/data/jsons/filenames"
)

func TestSignIn(t *testing.T) {
	userData := data.GetSignInMock(filenames.SignInMock)
	requestBody := bytes.NewReader(userData)

	clientRequest, err := http.NewRequest("POST", "http://localhost:8180/api/user/signin", requestBody)

	if err != nil {
		t.Errorf("Erro ao criar a solicitação: %v", err)
		return
	}
	clientRequest.Header.Set("Content-Type", "application/json")

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Errorf("Erro ao fazer a solicitação: %v", err)
		return
	}
	defer serverResponse.Body.Close()

	if serverResponse.StatusCode != http.StatusOK {
		t.Errorf("Esperado código de status %d, mas obteve %d", http.StatusCreated, serverResponse.StatusCode)
		return
	}
}

func TestSignUp(t *testing.T) {
	t.Skip("Not implemented")
}
