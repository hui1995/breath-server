FROM golang
LABEL maintainer=rong<breathcoder@gmail.com>

COPY . /$GOPATH/src/breath-server/
WORKDIR /$GOPATH/src/breath-server/

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct


RUN go build .

EXPOSE 9091
ENTRYPOINT ["./breath-server"]