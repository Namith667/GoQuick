FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o goquick .

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/goquick .

EXPOSE 8080

ENV PORT=8080 \
    LOG_LEVEL=INFO


CMD ["./goquick"]