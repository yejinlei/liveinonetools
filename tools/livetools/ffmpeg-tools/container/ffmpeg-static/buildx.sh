sudo docker buildx build --rm --platform linux/amd64,linux/arm64 -f Dockerfile.ffmpeg6  -t registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.ffmpeg.static:6.0 . --push
sudo docker buildx build --rm --platform linux/amd64,linux/arm64 -f Dockerfile.ffmpeg3  -t registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.ffmpeg.static:3.3.9 . --push
