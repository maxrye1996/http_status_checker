FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY status_checker/go.mod  ./
COPY status_checker/. .
RUN go mod download

RUN go build -o main .

# Use a minimal image for running the application
FROM alpine:latest

COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]