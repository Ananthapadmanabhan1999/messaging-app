FROM golang:1.21 as builder
 
WORKDIR /app
 
COPY go.mod go.sum ./
RUN go mod download
 
COPY websocket-service ./websocket-service
 
RUN CGO_ENABLED=0 GOOS=linux go build -o websocket-service ./websocket-service/cmd/main.go
 
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/websocket-service .
 
EXPOSE 8080
CMD ["./websocket-service"]