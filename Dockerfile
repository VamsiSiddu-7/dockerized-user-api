FROM golang:alpine

RUN mkdir /app

WORKDIR /app

RUN export GO111MODULE=on

COPY . /app/

RUN go build

EXPOSE 8080

ENTRYPOINT [ "./dockerized-user-api" ]
