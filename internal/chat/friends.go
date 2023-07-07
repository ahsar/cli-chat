package chat

import (
	"strconv"

	"github.com/eatmoreapple/openwechat"

	"log"
)

var (
	friendsMap []*openwechat.Friend

	// 内部id 转vxid
	friendsIdMap = make(map[string]int)
)

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
		name := GetName(fr.User)
		friendsMap[i] = fr
		friendsIdMap[fr.ID()] = i
		s = append(s, []string{strconv.Itoa(i), name})
	}

	return
}

func TalkToId(i int, s string) {
	u := friendsMap[i]
	if _, e := u.SendText(s); e != nil {
		log.Println("talk to id err", e)
	}
}

func FriendById(i int) *openwechat.Friend {
	return friendsMap[i]
}

// ConverVxid2Id
//
// conver vxid 2 local id
func ConverVxid2Id(i string) int {
	x, ok := friendsIdMap[i]
	if !ok {
		x = -1
	}
	return x
}

func GetName(u *openwechat.User) (s string) {
	s = u.NickName
	if u.RemarkName != "" {
		s = u.RemarkName
	}
	return
}
