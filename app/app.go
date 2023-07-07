package app

import (
	"log"

	"github.com/ahsar/cli-chat/internal/chat"
	"github.com/ahsar/cli-chat/internal/record"
	"github.com/ahsar/cli-chat/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eatmoreapple/openwechat"
)

func Run() {
	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		log.Fatal("log write err", err)
	}
	defer f.Close()

	// chat init
	// 接收微信通知消息
	ch := make(chan *openwechat.Message)
	chat.Init(ch)

	// record init
	record.NewRecord(100)

	// ui init
	if _, err := tea.NewProgram(
		ui.NewModel(ch), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("Error while running program:", err)
	}
	chat.Logout()
}
