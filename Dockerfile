FROM golang

WORKDIR /go/src/app
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o bin/www server.go

EXPOSE 8080
ENTRYPOINT ["./bin/www"]