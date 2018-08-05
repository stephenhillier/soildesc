FROM golang:1.10-alpine
RUN mkdir -p /go/src/github.com/stephenhillier/soildesc
ADD . /go/src/github.com/stephenhillier/soildesc/
RUN go install github.com/stephenhillier/soildesc/
ENTRYPOINT /go/bin/soildesc
EXPOSE 8000
USER 1001
