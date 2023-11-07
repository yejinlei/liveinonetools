sudo docker buildx build --platform linux/arm64,linux/amd64 -f Dockerfile . -t registry.cn-hangzhou.aliyuncs.com/yejinlei/vip-ingest:1.0.0 --push
