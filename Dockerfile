FROM golang:1.18.1-alpine3.15

RUN set -ex; \
    apk update; \
    apk add --no-cache git

RUN mkdir /app

WORKDIR /app

ADD . .

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
