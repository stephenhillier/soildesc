FROM golang:1.10-alpine AS builder
RUN mkdir -p /go/src/github.com/stephenhillier/soildesc
ADD . /go/src/github.com/stephenhillier/soildesc/
RUN go install github.com/stephenhillier/soildesc/

FROM alpine:3.8
WORKDIR /app
COPY --from=builder /go/bin/soildesc /app/
ENTRYPOINT /app/soildesc
EXPOSE 8000
USER 1001
