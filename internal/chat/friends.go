package chat

import (
	"fmt"

	"github.com/ahsar/cli-chat/internal/logger"
)

// Friends
// 获取账号下所有好友
func Friends() {
	friends, err := self.Friends()
	if err != nil {
		logger.Fatal("获取好友列表失败 %+v", err)
		return
	}

	for _, fr := range friends {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}
		fmt.Println(name)
	}
}
