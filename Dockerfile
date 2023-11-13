FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main ./bin/app

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main /app
RUN touch .env
CMD ["./main"]