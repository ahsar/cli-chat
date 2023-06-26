package chat

import (
	"strconv"

	"github.com/eatmoreapple/openwechat"

	"log"
)

var FriendsMap map[int]*openwechat.Friend

// Friends
// 获取账号下所有好友
// []string{id, name}
func Friends() (s [][]string) {
	friends, err := self.Friends()
	if err != nil {
		log.Fatal("获取好友列表失败", err)
		return
	}

	s = make([][]string, 0, len(friends))
	for i, fr := range friends {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}

		s = append(s, []string{strconv.Itoa(i), name})
	}

	return
}
