FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o linktok ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/linktok .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./linktok"] 