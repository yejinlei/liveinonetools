FROM registry.cn-hangzhou.aliyuncs.com/yejinlei/golang117.build:1.0 AS builder

ENV GIN_MODE=release
WORKDIR /go/src/app
COPY .  .

ARG TARGETOS TARGETARCH TARGETVARIANT

ENV CGO_ENABLED=0
RUN cd RTSPtoWebRTC \
    && go get \
    && go mod download \
    && GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#"v"} go build -a -o rtsp-to-webrtc

EXPOSE 8081

FROM ubuntu:focal AS runner
WORKDIR /app

COPY --from=builder /go/src/app/RTSPtoWebRTC/rtsp-to-webrtc /app/
COPY --from=builder /go/src/app/RTSPtoWebRTC/web /app/web

RUN mkdir -p /config
COPY --from=builder /go/src/app/RTSPtoWebRTC/config.json /app/

#ENV GO111MODULE="on"
#ENV GIN_MODE="release"

CMD ["./rtsp-to-webrtc"]
