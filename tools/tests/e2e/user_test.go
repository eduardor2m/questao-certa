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

	// Create a new request

	clientRequest, err := http.NewRequest("POST", baseURL+"/user/signin", requestBody)

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
	userData := data.GetSignUpMock(filenames.SignUpMock)
	requestBody := bytes.NewReader(userData)

	requestBody = bytes.NewReader(userData)

	clientRequest, err := http.NewRequest("POST", baseURL+"/user", requestBody)

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

func TestGetUser(t *testing.T) {
	t.Skip("Teste não implementado")
}

func TestUpdateUser(t *testing.T) {
	t.Skip("Teste não implementado")
}

func TestDeleteUser(t *testing.T) {
	t.Skip("Teste não implementado")
}

func TestGetUsers(t *testing.T) {
	t.Skip("Teste não implementado")
}
