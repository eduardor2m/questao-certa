FROM golang:1.21.0 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/application/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/api .

EXPOSE 8080

CMD ["./api"]
