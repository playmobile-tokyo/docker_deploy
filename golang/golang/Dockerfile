#FROM golang:1.10.3-alpine3.8
FROM golang:latest

CMD ["mkdir", "-pf", "/go/src"]

WORKDIR /go

ADD . /go

CMD ["go", "run", "./src/index.go"]

