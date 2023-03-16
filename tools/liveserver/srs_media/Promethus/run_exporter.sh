docker stop node-exporter
docker stop cadvisor
docker rm node-exporter
docker rm cadvisor

#端口9100
docker run --name node-exporter -d --privileged=true -v "/proc:/host/proc:ro" -v "/sys:/host/sys:ro" -v "/:/rootfs:ro" --network=host prom/node-exporter

#监控docker
docker run --name=cadvisor -d --privileged=true --volume=/:/rootfs:ro --volume=/var/run:/var/run:rw --volume=/sys:/sys:ro --volume=/var/lib/docker/:/var/lib/docker:ro --volume=/dev/disk/:/dev/disk:ro --publish=38080:8080 --detach=true google/cadvisor
