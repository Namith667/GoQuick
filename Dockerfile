FROM golang:alpine

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o goquick .

EXPOSE 8080

ENV PORT=8080 \
    LOG_LEVEL=INFO


CMD ["./goquick"]