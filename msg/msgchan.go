package msg

import (
	"github.com/byteweap/pkg/logs"
	"github.com/byteweap/pkg/redix"
)

// MsgChan 消息通道
type MsgChan struct {
	srcAppName    string       // 我的服务名 - 服务唯一标识
	targetAppName string       // 目标服务名 - 服务唯一标识
	rdx           *redix.Redix // 使用redis消息队列进行消息传输
}

func NewMsgChan(srcAppName, targetAppName string, rdx *redix.Redix) *MsgChan {
	return &MsgChan{
		srcAppName:    srcAppName,
		targetAppName: targetAppName,
		rdx:           rdx,
	}
}

// Push 推送消息给 targetAppName
func (gc *MsgChan) Push(data []byte) {
	key := Key(gc.srcAppName, MSG_TYPE_PUSH, gc.targetAppName)
	if err := gc.rdx.LPush(key, data); err != nil {
		logs.Error().Str("srcAppName", gc.srcAppName).Str("Act", MSG_TYPE_PUSH).Str("targetAppName", gc.targetAppName).Str("RedisKey", key).Err(err).Msg("向游戏通道写入数据失败!")
	}
}

// ListenPush 监听来自 targetAppName 的主推消息
func (gc *MsgChan) ListenPush(fn func(data []byte)) {
	key := Key(gc.targetAppName, MSG_TYPE_PUSH, gc.srcAppName)
	gc.rdx.ListenAppMsg(key, fn)
}

// Request 向targetAppName发送请求
func (gc *MsgChan) Request(data []byte) {
	key := Key(gc.srcAppName, MSG_TYPE_REQUEST, gc.targetAppName)
	if err := gc.rdx.LPush(key, data); err != nil {
		logs.Error().Str("srcAppName", gc.srcAppName).Str("Act", MSG_TYPE_REQUEST).Str("targetAppName", gc.targetAppName).Str("RedisKey", key).Err(err).Msg("向游戏通道写入数据失败!")
	}
}

// ListenRequest 监听来自 targetAppName 的请求
func (gc *MsgChan) ListenRequest(fn func(data []byte)) {
	key := Key(gc.targetAppName, MSG_TYPE_REQUEST, gc.srcAppName)
	gc.rdx.ListenAppMsg(key, fn)
}

// Response 向targetAppName发送响应
func (gc *MsgChan) Response(data []byte) {
	key := Key(gc.srcAppName, MSG_TYPE_RESPONSE, gc.targetAppName)
	if err := gc.rdx.LPush(key, data); err != nil {
		logs.Error().Str("srcAppName", gc.srcAppName).Str("Act", MSG_TYPE_REQUEST).Str("targetAppName", gc.targetAppName).Str("RedisKey", key).Err(err).Msg("向游戏通道写入数据失败!")
	}
}

// ListenResponse 监听来自 targetAppName 的响应
func (gc *MsgChan) ListenResponse(fn func(data []byte)) {
	key := Key(gc.targetAppName, MSG_TYPE_RESPONSE, gc.srcAppName)
	gc.rdx.ListenAppMsg(key, fn)
}
