sudo docker buildx build --platform linux/arm64 -f ./Dockerfile.multi -t registry.cn-hangzhou.aliyuncs.com/yejinlei/domain-ingest.armv8:1.0.0  . --load
