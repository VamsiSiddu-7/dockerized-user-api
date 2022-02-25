FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN cd /app && go build

ENTRYPOINT [ "/app/dockerized-user-api" ]
