# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o osto-cli ./cmd

# ---------- Runtime Stage ----------
FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    sqlite \
    libc6-compat

WORKDIR /app

COPY --from=builder /app/osto-cli .

RUN mkdir -p storage

CMD ["./osto-cli"]