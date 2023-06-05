package app

import (
	"flag"

	"fmt"
	"os"

	"github.com/ahsar/cli-chat/internal/logger"
	"github.com/ahsar/cli-chat/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
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
	//frList := "id1 user1\nid2 user2"

	if _, err := tea.NewProgram(
		ui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}
