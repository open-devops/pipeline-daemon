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

# Copy source
ADD . /go/src/github.com/open-devops/pipeline-daemon

# Build executable server daemon
RUN go get github.com/gorilla/mux \
 && go get gopkg.in/mgo.v2 \
 && go install github.com/open-devops/pipeline-daemon

# Run the Daemon Service by default when the container starts
ENTRYPOINT /go/bin/pipeline-daemon

# Service listens on port 8080.
EXPOSE 8080

# Volume shared with host server
VOLUME /go/bin