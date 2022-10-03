set -x
sudo docker stop ingest
sudo docker rm ingest
sudo docker run --rm -d --name ingest -v $(pwd)/argv.conf:/argv.conf:ro  -v $(pwd)/camera.conf:/camera.conf:ro registry.cn-hangzhou.aliyuncs.com/yejinlei/domain-ingest:1.0.0
