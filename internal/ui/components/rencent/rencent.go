// show rencent chat
package rencent

import (
	"log"

	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	textarea textarea.Model
	Focused  byte
}

func NewModel() Model {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = "rencent contacts"
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
	return Model{textarea: t}
}

func (m Model) Init() (t tea.Cmd) {
	return
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("rencent focus msg", msg)
	log.Println("rencent focus", m.textarea.Focused())
	log.Println("rencent value", m.textarea.Value())
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, tea.Batch(cmd)
}

func (m *Model) View() (s string) {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.textarea.View(),
	)
}

func (m *Model) SetSize(w, h int) {
	m.textarea.SetWidth(w)
	m.textarea.SetHeight(h)
}

func (m *Model) Focus() {
	m.textarea.Focus()
	m.Focused = constant.RencentPanel
}

func (m *Model) Blur() {
	m.textarea.Blur()
	m.Focused = 0
}
