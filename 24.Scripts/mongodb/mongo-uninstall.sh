#!/bin/bash
set -e

echo "🛑 停止 MongoDB 服务..."
sudo systemctl stop mongod || true
sudo systemctl disable mongod || true

echo "🧹 卸载 MongoDB 包..."
sudo apt-get purge -y mongodb-org* || true

echo "🧹 删除数据目录、日志目录和配置文件..."
sudo rm -rf /var/lib/mongodb
sudo rm -rf /var/log/mongodb
sudo rm -f /etc/mongod.conf

echo "🧹 清理残留 systemd 文件..."
sudo systemctl daemon-reload
sudo systemctl reset-failed

echo "✅ MongoDB 已完全卸载"