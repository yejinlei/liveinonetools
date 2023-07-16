sudo docker buildx build --rm --platform linux/amd64,linux/arm64 -f Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.ffmpeg.static:6.0 . --push
