// package chat provide vx msg dispacth
package chat

import (
	"github.com/eatmoreapple/openwechat"
)

// messageHandler
// vx 消息处理分发函数
func messageHandler(msg *openwechat.Message) {
	go func() {
		msgch <- msg
	}()
}
