FROM golang
WORKDIR /go/src/gitlab-ci
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cidemo .
ENTRYPOINT ["/go/src/gitlab-ci/cidemo"]
