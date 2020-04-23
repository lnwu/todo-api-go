FROM golang

WORKDIR /go/src/app
COPY . .

RUN go build hello.go

EXPOSE 8080
CMD ["hello"]