FROM golang:1.18.0-stretch

WORKDIR /app

COPY ./ /app

RUN go mod tidy

ENTRYPOINT [ "go", "run", "main.go" ]