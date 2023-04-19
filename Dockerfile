FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/davveo/lemonShop
COPY . $GOPATH/src/github.com/davveo/lemonShop
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./lemonShop"]
