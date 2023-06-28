package chat

import (
	"strconv"

	"github.com/eatmoreapple/openwechat"

	"log"
)

var friendsMap []*openwechat.Friend

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
	friendsMap = make([]*openwechat.Friend, l)

	for i, fr := range friends {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}

		friendsMap[i] = fr
		s = append(s, []string{strconv.Itoa(i), name})
	}

	return
}

func TalkToId(i int, s string) {
	u := friendsMap[i]
	u.SendText(s)
}
