FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /ja_golang_chat_service_v1
RUN mkdir -p /ja_golang_chat_service_v1/proto


WORKDIR /ja_golang_chat_service_v1


COPY ./proto/chat.pb.go /ja_golang_chat_service_v1/proto
COPY ./main.go /ja_golang_chat_service_v1

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o ja_run .

CMD ./ja_run
