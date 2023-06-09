package app

import (
	"log"

	"github.com/ahsar/cli-chat/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	// 登录
	//chat.Login()

	// 获取通讯录(friends)
	//frList := chat.Friends()
	//frList := "id1 user1\nid2 user2"
	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		log.Fatal("log write err", err)
	}
	defer f.Close()

	if _, err := tea.NewProgram(
		ui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("Error while running program:", err)
	}
}
