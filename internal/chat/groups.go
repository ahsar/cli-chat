package chat

import (
	"fmt"
	"strconv"

	"github.com/ahsar/cli-chat/internal/logger"
	"github.com/liushuochen/gotable"
)

func Groups() {
	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		logger.Fatal("获取群组列表失败 %+v", err)
		return
	}

	table, err := gotable.Create("id", "群组名")
	for i, fr := range groups {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}

		table.AddRow([]string{strconv.Itoa(i), name})
	}

	fmt.Println(table)
}