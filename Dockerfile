FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/air-verse/air@latest

EXPOSE 3000

CMD [ "air" ]
