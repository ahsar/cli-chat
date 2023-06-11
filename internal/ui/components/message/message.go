// Display for chat content
package message

import (
	"github.com/ahsar/cli-chat/internal/ui/components/dialog"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	dialog   dialog.Model
	textarea textarea.Model
}

func NewModel() Model {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = "message..."
	t.ShowLineNumbers = true
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

	return Model{
		textarea: t,
		dialog:   dialog.NewModel(),
	}
}

func (m Model) Init() (t tea.Cmd) {
	return
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, tea.Batch([]tea.Cmd{}...)
}

func (m *Model) SetSize(w, h int) {
	m.textarea.SetWidth(w / 3)
	m.textarea.SetHeight(h/2 - constant.HelpHeight - 1)

	m.dialog.SetSize(w/3, h/2-constant.HelpHeight-3)
}

func (m *Model) View() (s string) {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.textarea.View(),
		m.dialog.View(),
	)
}
