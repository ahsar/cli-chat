package chat

import (
	"strconv"

	"github.com/eatmoreapple/openwechat"

	"github.com/ahsar/cli-chat/internal/logger"
	ichat "github.com/ahsar/cli-chat/internal/ui/chat"
	itable "github.com/ahsar/cli-chat/internal/ui/table"
	"github.com/charmbracelet/bubbles/table"
)

type (
	Key struct {
	}

	dialog struct {
	}
)

func (k Key) OnEnter(t table.Row) {
	i, _ := strconv.Atoi(t[0])
	ichat.Chat(friendsMap[i], dialog{})
}

func (d dialog) OnEnter(o any, s string) {
	var (
		f  *openwechat.Friend
		ok bool
	)
	if f, ok = o.(*openwechat.Friend); !ok {
		logger.Error("dialog to friend err")
		return
	}

	self.SendTextToFriend(f, s)
}

var friendsMap map[int]*openwechat.Friend

// Friends
// 获取账号下所有好友
func Friends() {
	friends, err := self.Friends()
	if err != nil {
		logger.Fatal("获取好友列表失败 %+v", err)
		return
	}

	l := len(friends)
	rows := make([]table.Row, 0, l)
	friendsMap = make(map[int]*openwechat.Friend, l)
	for i, fr := range friends {
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}

		friendsMap[i] = fr
		rows = append(rows, []string{strconv.Itoa(i), name})
	}

	itable.Draw([]table.Column{
		{Title: "ID", Width: 4},
		{Title: "昵称", Width: 8},
	}, rows, Key{})
}
