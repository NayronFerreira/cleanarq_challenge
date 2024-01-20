# Dockerfile
# Stage 1: Build
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

# Stage 2: Run
FROM scratch

COPY --from=builder /app/main /main

EXPOSE 8080

CMD ["/main"]docker-compose upsudo lsof -i :15672