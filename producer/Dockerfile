FROM golang:1.12.4-alpine AS builder

ENV GODEBUG netdns=cgo

ADD ./ /go/src/github.com/kaplanmaxe/maxwell-kafka-go/producer
      
RUN apk add --no-cache --update alpine-sdk make && \
  cd /go/src/github.com/kaplanmaxe/maxwell-kafka-go/producer && \
  make install && \
  mv /go/bin/producer /usr/bin/

# Create the second stage with the most basic that we need - a 
# busybox which contains some tiny utilities like `ls`, `cp`, 
# etc. When we do this we'll end up dropping any previous 
# stages (defined as `FROM <some_image> as <some_name>`) 
# allowing us to start with a fat build image and end up with 
# a very small runtime image. Another common option is using 
# `alpine` so that the end image also has a package manager.
FROM alpine as final

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/producer /bin/

EXPOSE 8080

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "producer" ]