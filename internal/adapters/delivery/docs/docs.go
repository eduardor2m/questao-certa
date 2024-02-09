// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Eduardo Melo",
            "email": "deveduardomelo@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/question": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Cria uma questão de múltipla escolha, onde só é possível uma resposta correta",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Cria uma questão de múltipla escolha",
                "parameters": [
                    {
                        "description": "Dados da questão de múltipla escolha",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.QuestionDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Questão criada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/response.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Erro ao criar questão",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Deleta todas as questões",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Deleta todas as questões",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/question/filter": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Lista questões por filtro",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Lista questões por filtro",
                "parameters": [
                    {
                        "description": "Filtro para busca de questões",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.FilterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Question"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/question/import": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Importa questões de múltipla escolha a partir de um arquivo CSV",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Importa questões de múltipla escolha",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Arquivo CSV com as questões de múltipla escolha",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/question/{id}": {
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Deleta uma questão",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Deleta uma questão",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID da questão",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/question/{page}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Lista todas as questões",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Lista todas as questões",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Número da página",
                        "name": "page",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Question"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Cria um usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Cria um usuário",
                "parameters": [
                    {
                        "description": "Dados do usuário",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Usuário criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/response.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Erro ao criar usuário",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "Autentica um usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Autentica um usuário",
                "parameters": [
                    {
                        "description": "Dados do usuário",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/verify": {
            "get": {
                "description": "Verifica se o usuário está logado ou é admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Verifica se o usuário está logado ou é admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token de autenticação",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.FilterDTO": {
            "type": "object",
            "properties": {
                "discipline": {
                    "type": "string",
                    "example": "Engenharia de Produção"
                },
                "organization": {
                    "type": "string",
                    "example": "CESGRANRIO"
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                },
                "topic": {
                    "type": "string",
                    "example": "Administração da Produção"
                },
                "year": {
                    "type": "string",
                    "example": "2023"
                }
            }
        },
        "request.QuestionDTO": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "Aumentar a qualidade"
                },
                "discipline": {
                    "type": "string",
                    "example": "Engehnaria de Produção"
                },
                "model": {
                    "type": "string",
                    "example": "multiple_choice"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "[\"Aumentar a produtividade\"",
                        " \"Diminuir a produtividade\"",
                        " \"Aumentar a qualidade\"",
                        " \"Diminuir a qualidade\"]"
                    ]
                },
                "organization": {
                    "type": "string",
                    "example": "CESGRANRIO"
                },
                "question": {
                    "type": "string",
                    "example": "Qual o objetivo da administração da produção?"
                },
                "topic": {
                    "type": "string",
                    "example": "Administração da Produção"
                },
                "year": {
                    "type": "string",
                    "example": "2023"
                }
            }
        },
        "request.UserDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "dudu@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "Eduardo Melo"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "request.UserLoginDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "dudu@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.InfoResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.Question": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "discipline": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "organization": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                },
                "year": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Questão Certa API",
	Description:      "API para gerenciamento de questões e respostas",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
