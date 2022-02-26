FROM golang:latest

RUN mkdir /app

WORKDIR /app

RUN export GO111MODULE=on

RUN go install github.com/vamsikrishnasiddu/dockerized-user-api@latest

RUN cd /app && git clone https://github.com/vamsikrishnasiddu/dockerized-user-api.git

RUN cd /app/dockerized-user-api && go build

EXPOSE 8080

ENTRYPOINT [ "/app/dockerized-user-api/dockerized-user-api" ]
