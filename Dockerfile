FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./
COPY entrypoint.sh .

RUN go mod download
RUN go install github.com/air-verse/air@v1.61.7
RUN go install github.com/jackc/tern/v2@v2.3.2

EXPOSE 3000

ENTRYPOINT ["sh", "./entrypoint.sh"]
