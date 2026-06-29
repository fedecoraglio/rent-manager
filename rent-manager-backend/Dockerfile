FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o rent-manager ./cmd/api

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/rent-manager .

EXPOSE 8080

CMD ["./rent-manager"]