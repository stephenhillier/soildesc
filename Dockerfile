FROM golang:1.10-alpine
ADD . /go/src/github.com/stephenhillier/soildesc/
RUN go install github.com/stephenhillier/soildesc/
ENTRYPOINT /go/bin/soildesc
EXPOSE 8000
