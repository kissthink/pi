FROM golang:1.7

RUN mkdir -p /go/src/github.com/smhouse/pi
WORKDIR /go/src/github.com/smhouse/pi
COPY ./ /go/src/github.com/smhouse/pi/

RUN go get && go build
