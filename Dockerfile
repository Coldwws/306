FROM golang:1.25.9 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .
WORKDIR /app/cmd/
RUN CGO_ENABLED=0  GOOS=linux go build -o ../app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 5050
CMD ["./app"]