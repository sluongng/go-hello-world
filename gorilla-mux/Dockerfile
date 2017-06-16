FROM golang:latest
MAINTAINER sluongng@gmail.com

ENV GOPATH /opt/go:$GOPATH
ENV PATH /opt/go/bin:$PATH

WORKDIR /opt/go/src/app
COPY . .

RUN go-wrapper download
RUN go get github.com/derekparker/delve/cmd/dlv
RUN go-wrapper install

RUN go build -o main .
CMD ["./main"]