FROM golang:latest

WORKDIR /code

COPY . .

RUN go build -o main

CMD [ "main" ]