# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o app ./services/inventory/cmd/main.go


RUN chmod +x /app/app

# build a tiny docker image
FROM alpine:latest

RUN apk update && apk add --no-cache tzdata

ENV TZDIR=/usr/share/zoneinfo

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/app /app

CMD [ "/app/app" ]
