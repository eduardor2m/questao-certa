package utils

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/eduardor2m/questao-certa/tools/tests/data"
	"github.com/eduardor2m/questao-certa/tools/tests/data/jsons/filenames"
)

func GenerateToken() (*string, error) {
	userData := data.GetSignInMock(filenames.SignInMock)
	requestBody := bytes.NewReader(userData)

	clientRequest, err := http.NewRequest("POST", "http://questao-certa-air:8080/api/user/signin", requestBody)

	if err != nil {
		return nil, err
	}
	clientRequest.Header.Set("Content-Type", "application/json")

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		return nil, err
	}
	defer serverResponse.Body.Close()

	if serverResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("esperado código de status %d, mas obteve %d", http.StatusCreated, serverResponse.StatusCode)
	}

	token := serverResponse.Header.Get("Authorization")

	return &token, nil
}
