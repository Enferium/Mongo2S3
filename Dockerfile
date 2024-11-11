FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o mongo_backup

FROM alpine:latest

#install mongodb tools for mongodump
RUN apk add --no-cache mongodb-tools

WORKDIR /root/

COPY --from=builder /app/mongo_backup .

CMD ["./mongo_backup"]