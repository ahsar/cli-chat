// show rencent chat
package rencent

import (
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func (m *Model) View() (s string) {
	t := textarea.New()
	t.Prompt = ""
	//t.ShowLineNumbers = false
	t.Cursor.Style = constant.CursorStyle
	t.FocusedStyle.Placeholder = constant.FocusedPlaceholderStyle
	t.BlurredStyle.Placeholder = constant.PlaceholderStyle
	t.FocusedStyle.CursorLine = constant.CursorLineStyle
	t.FocusedStyle.Base = constant.FocusedBorderStyle
	t.BlurredStyle.Base = constant.BlurredBorderStyle
	t.FocusedStyle.EndOfBuffer = constant.EndOfBufferStyle
	t.BlurredStyle.EndOfBuffer = constant.EndOfBufferStyle
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	t.SetHeight(10)

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		//t.View(),
		t.View(),
	)
	//return lipgloss.NewStyle().
	////BorderTop(true).
	//BorderStyle(lipgloss.NormalBorder()).
	////MarginTop(1).
	//Render(
	////lipgloss.JoinHorizontal(
	////lipgloss.Left,
	////t.View(),
	////t.View(),
	////),
	//"\n\n",
	//)
}