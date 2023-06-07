package chat

import (
	"fmt"
	"strings"

	"github.com/eatmoreapple/openwechat"

	"log"
)

var FriendsMap map[int]*openwechat.Friend

// Friends
// 获取账号下所有好友
func Friends() (s string) {
	friends, err := self.Friends()
	if err != nil {
		log.Fatal("获取好友列表失败", err)
		return
	}

	var builder strings.Builder
	l := len(friends)
	FriendsMap = make(map[int]*openwechat.Friend, l)
	for i, fr := range friends {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}

		builder.WriteString(fmt.Sprintf("%d %s\n", i, name))
		FriendsMap[i] = fr
	}
	return builder.String()
}
