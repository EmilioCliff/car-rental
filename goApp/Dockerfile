FROM golang:1.22.3-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o goApp ./main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/goApp .
COPY --from=builder /app/static /app/static

EXPOSE 8080
CMD ["./goApp"]