FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ../../../ /app/

WORKDIR /app/server/cmd/bootstrap

RUN go build -o narvis-bootstrap bootstrap.go

FROM alpine:latest

COPY --from=builder /app/server/cmd/bootstrap/narvis-bootstrap /usr/local/bin/narvis-bootstrap

CMD ["narvis-bootstrap"]