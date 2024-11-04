FROM golang:1.23.0-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY service-a/ ./service-a
COPY service-b/ ./service-b

RUN go build -o service-a/service-a ./service-a/main.go
RUN go build -o service-b/service-b ./service-b/main.go

FROM alpine:3.18
WORKDIR /app

ARG SERVICE
COPY --from=builder /app/${SERVICE}/${SERVICE} /app/${SERVICE}

EXPOSE 8080
CMD ["/app/service-a"]  # Esse CMD será substituído no Docker Compose
