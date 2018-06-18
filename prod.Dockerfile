FROM golang:1.8.5-jessie

ENV PORT 3000

WORKDIR /go/src/github.com/skygate/skylabs-js

RUN go get github.com/golang/dep/cmd/dep

ADD . .
RUN dep ensure --vendor-only
RUN go install -v ./...

CMD ["skylabs-js"]
