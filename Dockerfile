FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Copy the .env file
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

# Stage 2: Run
FROM scratch

COPY --from=builder /app/main /main
COPY --from=builder /app/.env ./

EXPOSE 8080

CMD ["/main"]