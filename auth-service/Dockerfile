
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY auth-service ./auth-service

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./auth-service/cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/auth-service .

EXPOSE 8081
CMD ["./auth-service"]