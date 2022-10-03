FROM golang:latest as build
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
WORKDIR /go/cache
ADD build.sh .
RUN ./build.sh
 
FROM ubuntu:18.04 as prod
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/cache/RTSPtoWebRTC /RTSPtoWebRTC
WORKDIR /RTSPtoWebRTC
CMD ['./RTSPtoWebRTC']
