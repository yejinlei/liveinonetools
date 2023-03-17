#添加Prometheus的DataSource
curl -s -H "Content-Type: application/json" \
    -XPOST http://admin:admin@localhost:3000/api/datasources \
    -d '{
    "name": "srs-prometheus",
    "type": "prometheus",
    "access": "proxy", "isDefault": true,
    "url": "http://localhost:9090"
}'

#导入图表
curl https://raw.githubusercontent.com/ossrs/srs-grafana/main/dashboards/srs-stat-import.json | \
curl -s -H "Content-Type: application/json" \
    -XPOST http://admin:admin@localhost:3000/api/dashboards/db \
    --data-binary @-
