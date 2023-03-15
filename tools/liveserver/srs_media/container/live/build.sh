#./build.sh '5.0release' 
sudo docker build -f Dockerfile . -t registry.cn-hangzhou.aliyuncs.com/yejinlei/domain-live:$1
