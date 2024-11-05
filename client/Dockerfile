FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ../ /app/

WORKDIR /app/client

RUN go build -o narvis-client main.go

FROM alpine:latest

COPY --from=builder /app/client/narvis-client /usr/local/bin/narvis-client

CMD ["narvis-client"]