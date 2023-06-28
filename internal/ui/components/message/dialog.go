// user input box
package message

import (
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

var Dialog *DialogModel

type DialogModel struct {
	textarea textarea.Model
}

func NewDialogModel() *DialogModel {
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
	Dialog = &DialogModel{textarea: t}
	return Dialog
}

func (m DialogModel) Init() (t tea.Cmd) {
	return
}

func (m *DialogModel) Update(msg tea.Msg) (*DialogModel, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, tea.Batch(cmd)
}

func (m *DialogModel) SetSize(w, h int) {
	m.textarea.SetWidth(w)
	m.textarea.SetHeight(h)
}

func (m *DialogModel) View() (s string) {
	return m.textarea.View()
}

func (m *DialogModel) Focus() {
	m.textarea.Focus()
}

func (m *DialogModel) Blur() {
	m.textarea.Blur()
}
