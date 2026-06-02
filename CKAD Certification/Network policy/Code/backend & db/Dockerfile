FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o backend .

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["./backend"]
