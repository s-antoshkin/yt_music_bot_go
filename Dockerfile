FROM golang:1.19.4-alpine3.17 AS builder

WORKDIR /yt_music_bot_go

COPY . .

RUN go mod download

RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /yt_music_bot_go/bin/bot .

CMD [ "./bot" ]
