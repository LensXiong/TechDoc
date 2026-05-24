#!/bin/bash
set -e

echo "📦 添加 MongoDB 官方仓库..."
wget -qO - https://www.mongodb.org/static/pgp/server-7.0.asc | sudo gpg --dearmor -o /usr/share/keyrings/mongodb-archive-keyring.gpg

echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-archive-keyring.gpg ] https://repo.mongodb.org/apt/ubuntu $(lsb_release -cs)/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list

echo "🔄 更新包列表..."
sudo apt-get update

echo "📦 安装 MongoDB..."
sudo apt-get install -y mongodb-org

echo "🛠  启动 MongoDB 服务..."
sudo systemctl enable mongod
sudo systemctl start mongod

echo "✅ MongoDB 安装完成"
sudo systemctl status mongod --no-page