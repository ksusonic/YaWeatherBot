FROM golang:1.19-alpine

WORKDIR /app
ADD . /app/

RUN go mod download
RUN go build -o main cmd/main.go

ENTRYPOINT /app/main
