FROM golang:1.14
ADD . /go/src/app
WORKDIR /go/src/app
EXPOSE 80
CMD [ "go", "run", "main.go" ]

