FROM golang:1.23-alpine AS builder  

WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV GOARCH=amd64 GOOS=linux

RUN go build -o app src/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /api/app .
RUN chmod +x /root/app  # Garantir permissões de execução
EXPOSE 8080
CMD ["./app"]
