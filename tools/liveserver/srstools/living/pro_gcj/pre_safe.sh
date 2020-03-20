#二进制混淆
upx ./ingest
./build_docker_ingest_1.0.sh

#去除符号
strip ./gcj
upx ./live
./build_docker_live_1.2.sh

#去除镜像历史
#docker export xxxx | docker import - yejinlei/gcj_ingest:1.0
#docker export xxxx | docker import - yejinlei/gcj_live:1.2
