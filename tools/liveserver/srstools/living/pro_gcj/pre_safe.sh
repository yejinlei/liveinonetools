upx ./ingest_safe
docker export xxxx | docker import - yejinlei/live_ingest_safe.go:1.0
upx ./live_safe
docker export xxxx | docker import - yejinlei/srsv2.0-r6_centos7.6_safe:1.2
