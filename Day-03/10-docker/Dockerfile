FROM golang:latest

WORKDIR $GOPATH/src/my-app

copy . .

RUN go mod download

RUN go build .

EXPOSE 8080

CMD [./my-app]