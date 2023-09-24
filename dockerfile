FROM golang:1.21

WORKDIR /home/go/app

RUN apt-get update && apt-get install -y openssl

EXPOSE 8080

CMD [ "tail", "-f", "/dev/null" ]