FROM golang
WORKDIR /go/src/gitlab-ci
COPY ./ /go/src/gitlab-ci
RUN cd /go/src/gitlab-ci && \
    GOOS=linux GOARCH=amd64 go build -o myservice
ENTRYPOINT ["/go/src/gitlab-ci/myservice"]
