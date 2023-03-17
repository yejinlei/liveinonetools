#docker stop node-exporter
#docker stop cadvisor
docker stop pushgateway
docker stop prometheus
docker stop grafana
docker rm prometheus
docker rm pushgateway
#docker rm cadvisor
#docker rm node-exporter
docker rm grafana

#端口9100
#docker run --name node-exporter -d --privileged=true -v "/proc:/host/proc:ro" -v "/sys:/host/sys:ro" -v "/:/rootfs:ro" --network=host prom/node-exporter

#监控docker
#docker run --name=cadvisor -d --privileged=true --volume=/:/rootfs:ro --volume=/var/run:/var/run:rw --volume=/sys:/sys:ro --volume=/var/lib/docker/:/var/lib/docker:ro --volume=/dev/disk/:/dev/disk:ro --publish=38080:8080 --detach=true google/cadvisor

#端口9091
docker run --name=pushgateway --network=host -d prom/pushgateway

cat > prometheus.yml <<EOF
global:
  scrape_interval: 15s
  scrape_timeout: 10s
  evaluation_interval: 15s

scrape_configs:
  - job_name: "node"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.30:9100"]

  - job_name: "fastdfs"
    scrape_interval: 150s
    scrape_timeout: 100s
    evaluation_interval: 150s
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.70:9018"]

  - job_name: "cadvisor"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.30:38080"]

  - job_name: "live_forward"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.30:39972"]

  - job_name: "live_rtmp"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.30:39976"]

  - job_name: "live_hls"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.30:39977"]

  - job_name: "live_httpflv"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.30:39978"]

  - job_name: "live_dvr"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.10.10.30:39975"]
EOF

#端口9090
#docker run --rm -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 9090:9090 prom/prometheus
docker run --name prometheus  --rm -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml --network=host prom/prometheus

#端口3000
docker run --name grafana --rm -d -it --network=host -e GF_SECURITY_ADMIN_USER=admin -e GF_SECURITY_ADMIN_PASSWORD=admin -e GF_USERS_DEFAULT_THEME=light grafana/grafana
