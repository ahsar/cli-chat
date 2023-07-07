package rencent

import (
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableModel struct {
	textarea textarea.Model
	focused  byte
}

func NewTable() *tableModel {
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
	tx := &tableModel{textarea: t}
	Obj = tx
	return tx
}

func (m *tableModel) Init() (t tea.Cmd) {
	return
}

func (m *tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m *tableModel) View() (s string) {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.textarea.View(),
	)
}

func (m *tableModel) SetSize(w, h int) {
	m.textarea.SetWidth(w)
	m.textarea.SetHeight(h)
}

func (m *tableModel) Focus() {
	m.textarea.Focus()
	m.focused = constant.RencentPanel
}

func (m *tableModel) Blur() {
	m.textarea.Blur()
	m.focused = 0
}

func (m *tableModel) Focused() byte {
	return m.focused
}

func (m *tableModel) SetItems(any) {
}

func (m *tableModel) AddUser(i, j string) {
}
func (m *tableModel) SetUserMsg(i, j string) {
}
