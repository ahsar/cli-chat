package app

import (
	"log"

	//"github.com/ahsar/cli-chat/internal/chat"
	"github.com/ahsar/cli-chat/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	// 登录
	//chat.Login()

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
