#sudo docker buildx build --platform linux/arm64 -f ./Dockerfile -t registry.cn-hangzhou.aliyuncs.com/yejinlei/fastdfs.armv8:V5.11 . --load
sudo docker buildx build --platform linux/amd64 --build-arg MODEL="Release" -t registry.cn-hangzhou.aliyuncs.com/yejinlei/ZLMediaKit:6.0 . --load
