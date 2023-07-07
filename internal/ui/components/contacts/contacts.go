// Package contacts show user's contacts
package contacts

import (
	"github.com/ahsar/cli-chat/internal/ui/constant"

	"github.com/ahsar/cli-chat/internal/ui/components/message"
	"github.com/ahsar/cli-chat/internal/ui/components/rencent"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	enter key.Binding
}

type Model struct {
	table   table.Model
	keymap  keymap
	Focused byte
}

func NewModel() Model {
	return Model{
		table: setTable(),
		keymap: keymap{
			enter: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "select frients"),
			),
		},
	}
}

func (m *Model) Init() (t tea.Cmd) {
	return
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.enter):
			// 1. register message user
			message.Msg.SetUser(m.table.SelectedRow()[0])

			// 2. add to rencent contacts
			rencent.Obj.AddUser(m.table.SelectedRow()[0], "")

			// 3. blur current panel
			m.Blur()
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *Model) Focus() {
	m.table.Focus()
	m.Focused = constant.ContactPanel
}

func (m *Model) Blur() {
	m.table.Blur()
	m.Focused = 0
}

func (m *Model) View() (s string) {
	return m.table.View()
}

func setTable() (t table.Model) {
	columns := []table.Column{
		{Title: "id", Width: 10},
		{Title: "昵称", Width: 10},
	}

	t = table.New(
		table.WithColumns(columns),
		//table.WithFocused(true),
		table.WithHeight(16),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)
	//t.Focus()
	return
}

func (m *Model) SetRow(r []table.Row) {
	m.table.SetRows(r)
}

func (m *Model) SetSize(w, h int) {
	columns := []table.Column{
		{Title: "id", Width: w / 2},
		{Title: "昵称", Width: w / 2},
	}
	m.table.SetColumns(columns)
	m.table.SetWidth(w)
	m.table.SetHeight(h)
}
