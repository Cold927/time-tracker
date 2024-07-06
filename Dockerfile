FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o server ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app .

ENV APP_ENV docker
ENV PORT ${Port}

CMD [ "./server" ]