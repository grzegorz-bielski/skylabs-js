FROM golang:1.8.5-jessie

ENV PORT 3000

WORKDIR /go/src/github.com/skygate/skylabs-js/backend

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/pilu/fresh

ADD . .
RUN dep ensure --vendor-only

CMD ["fresh"]
