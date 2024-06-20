FROM golang:1.22.4-alpine3.19 AS builder
COPY . /app
WORKDIR /app
RUN go build -o main .

FROM alpine:3.19
ENV DOCKERIZE_VERSION v0.7.0

RUN apk update --no-cache \
  && apk add --no-cache wget openssl \
  && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin
COPY --from=builder /app/main /usr/local/bin/migrate-users
CMD /usr/local/bin/dockerize -timeout 20s -wait tcp://${DB_HOST}:${DB_PORT} /usr/local/bin/migrate-users