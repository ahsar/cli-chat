package app

import (
	"flag"

	"github.com/ahsar/cli-chat/internal/logger"
	"github.com/ahsar/cli-chat/internal/ui"
)

// verbose 模式
var v bool

func Init() {
	flag.BoolVar(&v, "v", false, "verbose 模式")
	flag.Parse()

	level := "info"
	if v {
		level = "debug"
	}
	logger.Init(level)
}

func Run() {
	// 登录
	//chat.Login()

	// 获取通讯录(friends)
	//frList := chat.Friends()
	frList := "id1 user1\nid2 user2"
	page := ui.New()
	page.SetContacts(frList)
	page.Run()
	page.G.Close()
}
