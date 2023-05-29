package chat

import (
	"fmt"

	"github.com/ahsar/cli-chat/internal/logger"
)

func Groups() {
	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		logger.Fatal("获取群组列表失败 %+v", err)
		return
	}

	for _, fr := range groups {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}
		fmt.Println("", name)
	}
}
