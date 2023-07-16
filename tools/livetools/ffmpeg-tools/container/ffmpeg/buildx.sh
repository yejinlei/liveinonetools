sudo docker buildx build --rm --platform linux/amd64,linux/arm64 -f Dockerfile.multi.ffmpeg6.deps  -t registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.ffmpeg.deps:6.0 . --push
