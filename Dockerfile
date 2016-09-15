# Pipeline Daemon Server Image for Open DevOps Pipeline
#
# VERSION : 1.0
#
FROM golang:alpine

MAINTAINER Open DevOps Team <open.devops@gmail.com>

ENV REFRESHED_AT 2016-09-09

# Install utility tools
RUN set -x \
    && apk add --no-cache git

# Build executable server daemon
RUN go get github.com/gorilla/mux \
 && go get gopkg.in/mgo.v2 \
 && go get github.com/open-devops/pipeline-daemon/...

# Service listens on port 8080
EXPOSE 8080

# Run the Daemon Service by default when the container starts
ENTRYPOINT /go/bin/pipeline-daemon