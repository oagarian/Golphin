FROM golang:latest

WORKDIR /api

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "src/main.go"]
