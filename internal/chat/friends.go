package chat

import (
	"strconv"

	"github.com/eatmoreapple/openwechat"

	"log"
)

var FriendsMap []*openwechat.Friend

// Friends
// 获取账号下所有好友
// []string{id, name}
func Friends() (s [][]string) {
	friends, err := self.Friends()
	if err != nil {
		log.Fatal("获取好友列表失败", err)
		return
	}

	l := len(friends)
	s = make([][]string, 0, l)
	FriendsMap = make([]*openwechat.Friend, l)

	for i, fr := range friends {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}

		FriendsMap[i] = fr
		s = append(s, []string{strconv.Itoa(i), name})
	}

	return
}
