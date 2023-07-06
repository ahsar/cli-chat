// show rencent chat
package rencent

import (
	tea "github.com/charmbracelet/bubbletea"
)

var Obj Rencent

type Rencent interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
	SetSize(int, int)
	Focus()
	Blur()
	Focused() byte
	SetItems(any)
	AddUser(string)
	SetUserMsg(i, j string)
}
