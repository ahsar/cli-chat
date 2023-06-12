// user input box
package dialog

import (
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textarea textarea.Model
}

func NewModel() Model {
	t := textarea.New()
	t.Prompt = ""
	t.ShowLineNumbers = true
	t.Placeholder = "Press ENTER to send"
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
	return Model{textarea: t}
}

func (m Model) Init() (t tea.Cmd) {
	return
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, tea.Batch(cmd)
}

func (m *Model) SetSize(w, h int) {
	m.textarea.SetWidth(w)
	m.textarea.SetHeight(h)
}

func (m *Model) View() (s string) {
	return m.textarea.View()
}

func (m *Model) Focus() {
	m.textarea.Focus()
}
