FROM golang:1.23.4
ADD . /go/src/app
WORKDIR /go/src/app
EXPOSE 9000
CMD [ "go", "run", "main.go" ]

