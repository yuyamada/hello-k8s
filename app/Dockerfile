FROM golang:latest
WORKDIR $GOPATH/src/github.com/yuyamada/hello-k8s/app
COPY . .
RUN go get -d -v ./...
ENTRYPOINT ["go", "run", "httpd/main.go"]
