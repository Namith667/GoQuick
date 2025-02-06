FROM golang:1.23 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download  # Download dependencies

COPY . .


RUN go build -o goquick

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/goquick .

EXPOSE 8080


ENV DB_HOST=DB \
    DB_USER=testuser \
    DB_PASS=1234 \
    DB_NAME=goquickdb \
    DB_PORT=5432 \
    DB_SSL_MODE=disable \
    PORT=8080 \
    JWT_SECRET_KEY=secret \
    JWT_EXPIRATION_TIME=72 \
    LOG_LEVEL=INFO 

CMD ["./goquick"]
