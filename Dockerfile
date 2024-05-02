FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp .

FROM alpine:latest

COPY --from=builder /app/myapp /usr/local/bin/

ENTRYPOINT ["myapp"]
