FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o app main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 CMD wget --no-verbose --tries=1 --spider http://localhost:8080/up || exit 1

CMD ["./app"]
