FROM golang:latest

ADD . /go/src/github.com/juanefec/minichat

WORKDIR /go/src/github.com/juanefec/minichat

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN go build -o main

EXPOSE 8080

CMD ["./main"]
