
# HTTP 请求并处理流式响应
功能：实现一个发送HTTP请求并处理服务器发送事件（SSE）流的功能，用于处理实时流数据的应用场景。
```
func doHttpStreamRequest(client *http.Client, method, link, body string, headers map[string]string) (reply chan string, code int, err error) {
	var reader io.Reader = nil
	if len(body) > 0 {
		reader = strings.NewReader(body)
	}

	req, err := http.NewRequest(method, link, reader)
	if err != nil {
		return
	}

	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Set(`User-Agent`, `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36`)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	code = resp.StatusCode

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream") {
		err = fmt.Errorf("not a event-stream response")
		a, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return
	}

	reply = make(chan string, 1)
	go func() {
		defer resp.Body.Close()
		defer func() {
			if errP := recover(); errP != nil {
				logs.ErrorF("doHttpStreamRequest.recover , err: %v", errP)
			}
		}()
		r := bufio.NewReader(resp.Body)
		for {
			line, readErr := r.ReadString('\n')
			if readErr != nil {
				break
			}
			if line != `` {
				reply <- line
			}
		}
		close(reply)
	}()
	return reply, code, nil
}
```


#  Goroutine 实践
实现功能：
使用 Go 代码实现一个 Goroutine，用于处理从 srcReply 通道接收到的消息，并将处理后的事件发送到 reply 通道。
代码的主要功能是解析从 srcReply 通道接收的消息，并根据特定的消息格式生成 types.SseReply 事件，最终将事件发送到 reply 通道。

数据格式：
```
{"content":" A","msg_id":"1796009433156104192","reply":"1796009432237551616","timestamp":1717036994277,"type":"reply"}
```

示例代码：
```
srcReply, code, err := util.SimpleHttpStream(http.MethodPost, xxxxx, map[string]string{
		`Content-Type`: `application/json`,
	}, string(d), util.DefaultHttpTimeout())

	reply = make(chan types.SseReply, 1)
	go func() {
		var evt = types.SseReply{}
		defer func() {
			if evt.Name != `` && evt.Data != `` {
				reply <- evt
			}
			close(reply)
		}()

		for {
			select {
			case <-ctxChat.Done():
				return
			case msg, ok := <-srcReply:
				fmt.Println(msg, ok)
				if !ok {
					return
				}
				if msg == "\n" {
					reply <- evt
					evt = types.SseReply{}
				}

				parts := strings.Split(msg, `: `)
				if len(parts) < 2 {
					continue
				}
				switch strings.TrimSpace(strings.ToLower(parts[0])) {
				case `id`:
					evt.Id = strings.TrimSpace(parts[1])
				case `event`:
					evt.Name = strings.TrimSpace(parts[1])
				case `data`:
					row := strings.Join(parts[1:], ": ")
					if row != "" {
						var rc RCResult
						json.Unmarshal([]byte(row), &rc)
						row = rc.Data.Text
						evt.RealContent = rc.Data.RealContent
					}
					evt.Data += row
				}
			}
		}
	}()
```
代码解释：
```
使用 for 循环和 select 语句等待多个通信操作。
case <-ctxChat.Done(): 如果接收到上下文 ctxChat 的 Done 信号，则退出 Goroutine。
case msg, ok := <-srcReply: 从 srcReply 通道接收消息。
打印消息和 ok 状态。
如果 ok 为 false（即通道关闭），则退出 Goroutine。
如果消息是换行符 "\n"，则将当前 evt 发送到 reply 通道，并重置 evt。
将消息按 ": " 分割成部分，如果长度小于 2，跳过该消息。
根据消息的前缀部分（例如 id, event, data）更新 evt 的相应字段。
如果前缀是 data，进一步处理数据部分，将其 JSON 反序列化并更新 evt 的 Data 和 RealContent 字段。
```

1、主要逻辑
* 处理上下文取消: 当 ctxChat 的上下文被取消时，Goroutine 退出。
* 接收和处理消息: 从 srcReply 通道接收消息并解析成事件。
* 解析消息: 根据消息的前缀（id, event, data）更新 evt 的字段。
* 发送事件: 在特定条件下将 evt 发送到 reply 通道，并在 Goroutine 结束时处理剩余的 evt。 

2、关键点：
* defer 确保在 Goroutine 结束时适当处理资源（发送最后一个事件并关闭通道）。
* 使用 select 语句处理多路复用通道操作。
* 根据特定的消息格式解析并构建事件对象 types.SseReply。
这段代码整体上实现了一个异步事件处理系统，能够从一个源通道接收消息，解析消息并生成事件，然后通过另一个通道将这些事件传递出去。


# 问题解答
