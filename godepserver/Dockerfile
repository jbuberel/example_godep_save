# Dockerfile extending the generic Go image with application files for a
# single application.
FROM gcr.io/google_appengine_qa/golang

COPY . /go/src/app
ENV GOPATH $GOPATH:/go/src/app/Godeps/_workspace
RUN go get github.com/kr/godep
RUN godep go build -o /go/bin/app
