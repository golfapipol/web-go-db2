FROM golang:1.15 as builder
RUN apt update && apt install -y libxml2 gcc build-essential
RUN go get -d github.com/ibmdb/go_ibm_db
RUN cd /go/src/github.com/ibmdb/go_ibm_db/installer && go run setup.go
ENV DB2HOME=/go/src/github.com/ibmdb/go_ibm_db/installer/clidriver \
CGO_CFLAGS=-I/go/src/github.com/ibmdb/go_ibm_db/installer/clidriver/include \
CGO_LDFLAGS=-L/go/src/github.com/ibmdb/go_ibm_db/installer/clidriver/lib \
LD_LIBRARY_PATH=/go/src/github.com/ibmdb/go_ibm_db/installer/clidriver/lib
WORKDIR /module
COPY ./go.mod /module/go.mod
COPY ./go.sum /module/go.sum
RUN go mod download
COPY . /module/
RUN GOOS=linux go build -o ./bin/app main.go

ENV GIN_MODE release
EXPOSE 4000
CMD ./app