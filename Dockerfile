FROM golang:1.14
ADD . /go/src/app
WORKDIR /go/src/app
EXPOSE 9203
CMD [ "go", "run", "main.go" ]

