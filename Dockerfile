FROM golang:buster

WORKDIR /go/src/app
COPY . .
RUN GOPATH=/go/src/app make

CMD cmd/bourbon-finder/bourbon-finder -c /go/src/app/data/config.json -w
