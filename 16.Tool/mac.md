# 两台 Mac 开启通用控制的完整步骤

原理：通过蓝牙、Wi-Fi 和 Handoff 技术，系统自动检测附近的设备，并同步鼠标键盘事件。

* 支持系统：macOS Monterey 12.3 及以上。
* 硬件要求：2016 年后的 Mac 绝大多数都支持。
* 连接要求：两台设备登录同一个 Apple ID，开启蓝牙、Wi-Fi，并且在系统偏好设置里开启“通用控制”。
* 使用体验：鼠标直接移到屏幕边缘就能“穿”到另一台 Mac，键盘自动跟随鼠标所在设备，拷贝/粘贴也能跨设备。

开启通用控制步骤：

* 确认系统版本：两台 Mac 都要 macOS Monterey 12.3 及以上（建议升级到最新版，修复了很多通用控制的 bug）。
  点击  → 关于本机 查看版本。
* 开启 Handoff 与蓝牙/Wi-Fi
  在两台 Mac 上： 打开 系统设置（或系统偏好设置）→ 通用 → 隔空投送与接力（旧版本在「通用」下的「接力」）
  打开 允许在这台 Mac 与 iCloud 设备之间使用接力，蓝牙、Wi-Fi 都必须开启，且连接到同一个局域网（同一个路由器的 2.4G 或 5G 都可以）。
* 开启通用控制：在两台 Mac 上： 打开 系统设置 → 显示器 → 通用控制…（旧版本在「显示器」中点「通用控制」按钮）
  勾选： 允许你的光标和键盘在任意 Mac 或 iPad 间移动 ，推动屏幕边缘连接另一台设备 ，自动重新连接到附近的 Mac 或 iPad
* 排列显示器：
  在 显示器 设置里，拖动屏幕位置（上方的排列图）来模拟物理摆放，比如左 Mac 在右 Mac 的左边，就把它放在左边。
  光标到达屏幕边缘时会“滑”到另一台设备。
* 测试
  移动鼠标到两台 Mac 相邻边缘，停留 1~2 秒，出现灰色边框动画，松手即可穿过去。
  键盘焦点会自动跟随鼠标设备。

拍错记录：
① 隔空投送需要改成任何人。


通用控制自动修复脚本：

```
#!/bin/bash

echo "==================== 通用控制自动修复脚本（Mac mini） ===================="

# 1️⃣ 开启接力 (Handoff)
echo "1️⃣ 正在开启 Handoff / 接力..."
defaults write com.apple.coreservices.useractivityd.plist ActivityAdvertisingAllowed -bool true
defaults write com.apple.coreservices.useractivityd.plist ActivityReceivingAllowed -bool true
echo "✅ 已启用 Handoff 接力功能"

# 2️⃣ 检查蓝牙状态
echo "2️⃣ 检查蓝牙状态..."
bt_status=$(system_profiler SPBluetoothDataType | grep "State:" | awk '{print $2}')
if [ "$bt_status" != "On" ]; then
  echo "⚠️ 蓝牙未开启，正在尝试开启..."
  sudo defaults write /Library/Preferences/com.apple.Bluetooth ControllerPowerState -int 1
  sudo killall -HUP blued
  echo "✅ 蓝牙已开启"
else
  echo "✅ 蓝牙正常"
fi

# 3️⃣ 检查 Wi-Fi 状态
echo "3️⃣ 检查 Wi-Fi 状态..."
wifi_status=$(networksetup -getairportpower en1 | awk '{print $4}')
if [ "$wifi_status" != "On" ]; then
  echo "⚠️ Wi-Fi 未开启，正在开启..."
  networksetup -setairportpower en1 on
else
  echo "✅ Wi-Fi 已开启"
fi

# 4️⃣ 检查并提示通用控制设置
echo "4️⃣ 检查通用控制设置..."
echo "⚙️ 请手动确认以下设置："
echo "系统设置 → 显示器 → 通用控制 → 勾选三项："
echo " - 允许光标和键盘在任意 Mac 或 iPad 间移动"
echo " - 推动屏幕边缘连接另一台设备"
echo " - 自动重新连接到附近的 Mac 或 iPad"

# 5️⃣ 重启相关进程
echo "5️⃣ 正在重启接力与通用控制相关服务..."
killall useractivityd 2>/dev/null
killall ControlCenter 2>/dev/null
killall cfprefsd 2>/dev/null
echo "✅ 已重启相关服务"

# 6️⃣ 检查网络连通性
echo "6️⃣ 测试网络连通性（可忽略 ping 超时）"
ping -c 2 10.17.0.75
ping -c 2 10.17.43.29

echo "==================== 修复完成 ===================="
echo "💡 建议操作："
echo "1️⃣ 打开 系统设置 → 通用 → 接力 确认已开启"
echo "2️⃣ 打开 系统设置 → 显示器 → 通用控制 勾选三项"
echo "3️⃣ 三台 Mac 都执行以上设置，然后重启蓝牙和 Wi-Fi"
echo "4️⃣ 移动鼠标至屏幕边缘测试通用控制连接"

```
