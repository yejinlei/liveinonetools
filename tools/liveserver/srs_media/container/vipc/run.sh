set -x
sudo docker stop vipc
sudo docker rm vipc
sudo docker run --rm -d --name vipc -p 8554:8554 -v $PWD/rtsp-simple-server.yml:/rtsp-simple-server/rtsp-simple-server.yml:ro registry.cn-hangzhou.aliyuncs.com/yejinlei/domain-vipc:1.0.0
