sudo docker buildx build --platform linux/arm64 -f ./Dockerfile.armv8 -t registry.cn-hangzhou.aliyuncs.com/yejinlei/domain-live.armv8:1.0.0  . --load
#sudo docker build -f ./Dockerfile.armv8 -t registry.cn-hangzhou.aliyuncs.com/yejinlei/domain-live.armv8:1.0.0  .
