// user input box
package dialog

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() (t tea.Cmd) {
	return
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, tea.Batch([]tea.Cmd{}...)
}

func (m Model) View() (s string) {
	return
}
