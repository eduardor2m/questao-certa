# Questão Certa - API Golang com Framework Echo e Banco de Dados MongoDB

Esta é a documentação para a API Questão Certa, uma aplicação Golang que utiliza o framework Echo para criar uma API RESTful para gerenciar questões de múltipla escolha com um banco de dados MongoDB.

## Endpoints

A API possui os seguintes endpoints:

- `/question`: Cria uma questão de múltipla escolha.
- `/question/import`: Importa questões de múltipla escolha a partir de um arquivo CSV.
- `/question/{page}`: Lista todas as questões de múltipla escolha.
- `/question/filter`: Lista questões de múltipla escolha com base em critérios de filtro.
- `/question/{id}`: Deleta uma questão de múltipla escolha com base no ID da questão.
- `/question`: Deleta todas as questões de múltipla escolha.

## Iniciar o Servidor

Para iniciar o servidor, utilize o seguinte comando:

```bash
go run ./cmd/application/main.go
```

Isso iniciará a API e a tornará acessível em `http://localhost:<PORT>`, onde `<PORT>` é a porta configurada no arquivo `.env`.

Isso conclui a documentação da API Questão Certa. Utilize os endpoints conforme necessário para gerenciar questões de múltipla escolha.
