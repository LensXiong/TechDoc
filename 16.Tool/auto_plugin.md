
# 微信一键双开或者多开终端脚本
`wechat-multi-open.sh`:
```
#!/bin/bash

# 你可以根据需要修改这个副本数量（默认只复制一个副本用于双开）
CLONE_COUNT=1

# 原始微信路径（请确保微信已安装）
SRC_APP="/Applications/WeChat.app"

# 检查原始微信是否存在
if [ ! -d "$SRC_APP" ]; then
    echo "❌ 原始 WeChat.app 未安装在 /Applications 目录下，脚本终止。"
    exit 1
fi

for i in $(seq 1 $CLONE_COUNT); do
    CLONE_APP="/Applications/wechat-${i}.app"

    echo "🌀 正在创建副本 wechat-${i}.app..."

    # 拷贝副本
    cp -R "$SRC_APP" "$CLONE_APP"

    # 修改 Bundle Identifier
    INFO_PLIST="${CLONE_APP}/Contents/Info.plist"
    if [ -f "$INFO_PLIST" ]; then
        /usr/libexec/PlistBuddy -c "Set :CFBundleIdentifier com.tencent.xinWeChat${i}" "$INFO_PLIST"
    fi

    # 解除隔离
    xattr -dr com.apple.quarantine "$CLONE_APP"

    # 重签名
    echo "🔏 重签名 wechat-${i}.app..."
    sudo codesign --force --deep --sign - "$CLONE_APP"

    # 启动副本
    echo "🚀 启动 wechat-${i}.app..."
    open -n "$CLONE_APP"
done

echo "✅ 微信副本已创建并启动完毕。你现在可以登录第二个账号啦！"
```
使用方式:

* 将上面的内容复制到 wechat-multi-open.sh
* 给脚本执行权限（在终端中运行）： `chmod +x wechat-multi-open.sh`
* 运行脚本（需要输入密码用于签名）： `./wechat-multi-open.sh`