FROM golang:1.17 as builder
ENV GOPROXY https://goproxy.cn
ENV GO11MODULE on
COPY . /
WORKDIR /rtsp-simple-server
RUN go build -mod=vendor -o rtsp-simple-server .

ENTRYPOINT /rtsp-simple-server/rtsp-simple-server
