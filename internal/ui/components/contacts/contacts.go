// user's contacts
package contacts

import (
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	table table.Model
}

func NewModel() Model {
	return Model{table: setTable()}
}

func (m Model) Init() (t tea.Cmd) {
	return
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	log.Println("contacts focus", m.table.Focused())
	log.Println("contacts Selected", m.table.SelectedRow())
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *Model) Focus() {
	m.table.Focus()
}

func (m *Model) Blur() {
	m.table.Blur()
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
	t.Focus()
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
