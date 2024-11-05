FROM golang:1.23-alpine as builder

WORKDIR /app

COPY ../ /app/

WORKDIR /app/server

RUN go build -o narvis-server main.go

FROM alpine:latest

COPY --from=builder /app/server/narvis-server /usr/local/bin/narvis-server

CMD ["narvis-server"]