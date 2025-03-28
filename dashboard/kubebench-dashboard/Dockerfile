FROM golang

ADD . /go/src/github.com/liayan/kubebench/dashboard
RUN go build github.com/liayan/kubebench/dashboard

ENTRYPOINT ./dashboard
EXPOSE 9303
