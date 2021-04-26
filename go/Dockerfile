FROM golang:latest

WORKDIR /go/app/

COPY . .

ENV PATH="/go/bin:${PATH}"

RUN apt-get update

RUN go get "github.com/go-sql-driver/mysql"

WORKDIR /go/app/src

CMD ["tail", "-f", "/dev/null"]