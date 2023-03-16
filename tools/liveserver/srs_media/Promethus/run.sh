docker stop node-exporter
docker stop prometheus
docker stop grafana
docker rm prometheus
docker rm node-exporter
docker rm grafana

#docker run --rm -d -p 9100:9100 prom/node-exporter
#docker run --name node-exporter --rm -d --network=host prom/node-exporter
docker run --name node-exporter -d  \
-v "/proc:/host/proc:ro" \
-v "/sys:/host/sys:ro" \
-v "/:/rootfs:ro" \
--net="host" \
prom/node-exporter

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
      - targets: ["10.13.2.11:9100"]
  - job_name: "live_forward"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.13.2.11:39972"]
  - job_name: "live_rtmp"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.13.2.11:39976"]
  - job_name: "live_hls"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.13.2.11:39977"]
  - job_name: "live_httpflv"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.13.2.11:39978"]
  - job_name: "live_dvr"
    metrics_path: "/metrics"
    scheme: http
    follow_redirects: true
    static_configs:
      - targets: ["10.13.2.11:39975"]
EOF
#docker run --rm -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 9090:9090 prom/prometheus
docker run --name prometheus  --rm -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml --network=host prom/prometheus

#docker run --rm -d -it -p 3000:3000 \
docker run --name grafana --rm -d -it --network=host -e GF_SECURITY_ADMIN_USER=admin -e GF_SECURITY_ADMIN_PASSWORD=admin -e GF_USERS_DEFAULT_THEME=light grafana/grafana
