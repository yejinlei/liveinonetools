FROM registry.cn-hangzhou.aliyuncs.com/yejinlei/golang117.build:1.0 AS builder

ENV GIN_MODE=release
WORKDIR /go/src/app
COPY .  .

ARG TARGETOS TARGETARCH TARGETVARIANT

ENV CGO_ENABLED=0
RUN cd RTSPtoWeb \
    && go get \
    && go mod download \
    && GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#"v"} go build -a -o rtsp-to-web

EXPOSE 8081

FROM ubuntu:focal AS runner
WORKDIR /app

COPY --from=builder /go/src/app/RTSPtoWeb/rtsp-to-web /app/
COPY --from=builder /go/src/app/RTSPtoWeb/web /app/web

RUN mkdir -p /config
COPY --from=builder /go/src/app/RTSPtoWeb/config.json /config

#ENV GO111MODULE="on"
#ENV GIN_MODE="release"

CMD ["./rtsp-to-web", "--config=/config/config.json"]
