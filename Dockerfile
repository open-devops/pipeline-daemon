#
# Pipeline Daemon Server Image for Open DevOps Pipeline
#
# VERSION : 1.0
#
FROM golang:alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go build -o daemonServe .

CMD ["/app/daemonServe"]