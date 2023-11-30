FROM bitnami/golang:latest

LABEL maintainer="adelahera@correo.ugr.es" \
  version="1.1"

RUN adduser --disabled-password -u 1001 basket-stats-test

USER basket-stats-test

WORKDIR /app

ENV GOPROXY=proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY go.mod go.sum taskfile.yaml ./

RUN go mod download

WORKDIR /app/test

ENTRYPOINT ["task", "test"]