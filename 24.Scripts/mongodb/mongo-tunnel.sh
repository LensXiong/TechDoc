#!/bin/bash
# mongo-tunnel.sh
# 自动建立本地到远程 MongoDB 的 SSH 隧道（支持密码或免密登录）
# 自动检测并安装 sshpass（仅密码登录时需要）

# 配置
REMOTE_USER="xxx"
REMOTE_HOST="xxxx"
REMOTE_PORT=27017
LOCAL_PORT=27017

echo "检测 sshpass..."
if ! command -v sshpass >/dev/null 2>&1; then
    read -p "sshpass 未安装，是否自动安装？(y/n): " install_sshp
    if [[ "$install_sshp" =~ ^[Yy]$ ]]; then
        if [[ "$(uname)" == "Darwin" ]]; then
            echo "MacOS 系统，使用 brew 安装 sshpass..."
            brew install hudochenkov/sshpass/sshpass
        else
            echo "Linux 系统，使用 apt 安装 sshpass..."
            sudo apt update
            sudo apt install -y sshpass
        fi
    else
        echo "请先安装 sshpass 后再使用密码登录方式"
    fi
fi

echo "选择 SSH 登录方式："
echo "1) 公钥免密登录"
echo "2) 输入密码登录"
read -p "请选择 1 或 2: " login_mode

if [[ "$login_mode" == "1" ]]; then
    echo "使用公钥免密方式连接..."
    ssh -fNL ${LOCAL_PORT}:127.0.0.1:${REMOTE_PORT} ${REMOTE_USER}@${REMOTE_HOST}
elif [[ "$login_mode" == "2" ]]; then
    if ! command -v sshpass >/dev/null 2>&1; then
        echo "sshpass 未安装，无法使用密码登录"
        exit 1
    fi
    read -s -p "请输入远程用户 ${REMOTE_USER} 密码: " REMOTE_PASS
    echo
    echo "使用密码方式连接..."
    sshpass -p "${REMOTE_PASS}" ssh -o StrictHostKeyChecking=no -fNL ${LOCAL_PORT}:127.0.0.1:${REMOTE_PORT} ${REMOTE_USER}@${REMOTE_HOST}
else
    echo "无效选择"
    exit 1
fi

echo "隧道已建立，本地端口 ${LOCAL_PORT} -> 远程 ${REMOTE_HOST}:${REMOTE_PORT}"
echo "你可以在本地执行："
echo "  mongosh --port ${LOCAL_PORT}"