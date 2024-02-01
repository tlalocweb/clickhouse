#!/bin/bash
mkdir -p ch_data
mkdir -p ch_logs
echo "** Starting Clickhouse for testing - will stay in foreground **" 
docker run --rm \
    -p 9942:9000 -p 8842:8123 \
    --cap-add=SYS_NICE --cap-add=NET_ADMIN --cap-add=IPC_LOCK \
    -e CLICKHOUSE_DB=gorm -e CLICKHOUSE_PASSWORD=gorm -e CLICKHOUSE_USER=gorm -e CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=1 \
    -v $(realpath ./ch_data):/var/lib/clickhouse/ \
    -v $(realpath ./ch_logs):/var/log/clickhouse-server/ \
    --name gorm-test-clickhouse --ulimit nofile=262144:262144 clickhouse/clickhouse-server
echo "** Cleaning up Clickhouse directories **"
rm -rf ch_data
rm -rf ch_logs
echo "Done."
