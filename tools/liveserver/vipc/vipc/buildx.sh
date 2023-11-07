#sudo docker build -f Dockerfile . -t registry.cn-hangzhou.aliyuncs.com/yejinlei/domain-vipc:1.0.0
sudo docker buildx build --platform linux/arm64,linux/amd64 -f Dockerfile.multi -t registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.vipc:1.0.0 . --push

