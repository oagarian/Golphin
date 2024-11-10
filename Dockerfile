FROM golang:latest AS builder
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod downloa
COPY . .
RUN go build -o app src/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /api/app .
EXPOSE 8080
CMD ["./app"]