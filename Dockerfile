FROM golang:1.23.4
ADD . /go/src/app
WORKDIR /go/src/app
EXPOSE 80
CMD [ "go", "run", "main.go" ]

