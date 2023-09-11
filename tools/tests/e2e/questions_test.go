package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateQuestion(t *testing.T) {
	jsonData := map[string]string{
		"question":     "Qual a cor do cavalo branco de Napoleão?",
		"answer":       "Branco",
		"options":      "Branco, Preto, Amarelo, Vermelho",
		"organization": "ENEM",
		"model":        "multiple_choice",
		"year":         "2019",
		"discipline":   "Matemática",
		"topic":        "Álgebra",
	}

	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", "/api/questions", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	assert.Equal(t, http.StatusCreated, resp.Code)
}
