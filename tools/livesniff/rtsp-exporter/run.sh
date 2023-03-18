#端口9018
sudo docker run --privileged=true -it --rm --network=host -v `pwd`/conf:/etc/media registry.cn-hangzhou.aliyuncs.com/yejinlei/media.exporter:v1.0.0
