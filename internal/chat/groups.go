package chat

import (
	"log"
)

func Groups() {
	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		log.Fatal("获取群组列表失败", err)
		return
	}

	for i, fr := range groups {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}

		log.Println("groups", i, name)
	}
}
