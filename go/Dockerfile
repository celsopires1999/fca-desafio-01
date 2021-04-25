FROM golang:latest

WORKDIR /go/app/src

ENV PATH="/go/bin:${PATH}"

RUN apt-get update

RUN go get "github.com/go-sql-driver/mysql"

CMD ["tail", "-f", "/dev/null"]