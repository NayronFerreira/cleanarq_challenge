# Dockerfile
FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

EXPOSE 8080

CMD ["/app/main"]