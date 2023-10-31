FROM golang:1.21

WORKDIR /home/go/app

RUN apt-get update && apt-get install -y openssl

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.1

RUN go install go.uber.org/mock/mockgen@v0.2.0

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.20.0

EXPOSE 8080

CMD [ "tail", "-f", "/dev/null" ]