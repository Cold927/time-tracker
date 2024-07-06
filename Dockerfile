FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/main.go

FROM debian:12.4-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/.env.docker /app/.env.docker
COPY --from=builder /app/docs app/docs

ENV APP_ENV docker
ENV PORT ${Port}

CMD [ "app/server" ]