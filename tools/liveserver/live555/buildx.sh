sudo docker buildx build --platform linux/amd64,linux/arm64 -f Dockerfile.multi.live555  -t registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.live555:latest . --push
