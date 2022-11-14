FROM golang:1.18.0-stretch

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go build -o /countries-api

EXPOSE 8080

CMD [ "countries-api"]