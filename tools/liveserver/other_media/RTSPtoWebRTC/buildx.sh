sudo docker buildx build --platform linux/arm64,linux/amd64 -f Dockerfile.multi -t registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.rtsp2webrtc:1.0.0 . --push
