package msg

import "fmt"

const (
	MSG_TYPE_REQUEST  = "req"  // 请求
	MSG_TYPE_RESPONSE = "resp" // 响应
	MSG_TYPE_PUSH     = "push" // 主推
)

// Key 构造消息队列的key
// desc: 使用:隔开, 第一个为发消息的服务名, 第二个为消息类型, 第三个为收消息的服务明
// 举例: a:req:b  -  a向b发送请求数据
// srcAppName: 发消息的appName
// msgType: 消息类型，请求: MSG_TYPE_REQUEST(req), 响应: MSG_TYPE_RESPONSE(resp)
// targetAppName: 收消息的appName
func Key(srcAppName, msgType, targetAppName string) string {
	return fmt.Sprintf("%s:%s:%s", srcAppName, msgType, targetAppName)
}

func MsgChanId(srcAppName, targetAppName string) string {
	return fmt.Sprintf("%s-%s", srcAppName, targetAppName)
}
