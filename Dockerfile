FROM golang:1.21-alpine3.17

LABEL maintainer="adelahera@correo.ugr.es" \
  version="1.1"

RUN adduser -D -u 1001 test

USER test

WORKDIR /app/test

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY go.mod go.sum taskfile.yaml ./

RUN go mod download

ENTRYPOINT ["task", "test"]