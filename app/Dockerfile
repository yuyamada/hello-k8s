FROM golang:latest
WORKDIR $GOPATH/src/github.com/yuyamada/hello-k8s/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
ENTRYPOINT app
