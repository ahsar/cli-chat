// Display for chat content
package message

import (
	"strings"

	"github.com/ahsar/cli-chat/internal/record"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var Msg *Model

type Model struct {
	dialog   *DialogModel
	textarea textarea.Model
	Focused  byte
}

func NewModel() *Model {
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

	Msg = &Model{
		textarea: t,
		dialog:   NewDialogModel(),
	}
	return Msg
}

func (m *Model) Init() (t tea.Cmd) {
	return
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.textarea, cmd = m.textarea.Update(msg)

	m.dialog, cmd = m.dialog.Update(msg)

	return m, tea.Batch(cmd)
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

func (m *Model) SetUser(id string) {
	m.Reset()

	m.dialog.SetUser(id)
	m.Focus()
}

func (m *Model) Focus() {
	m.Focused = constant.DialogPanel
	m.dialog.Focus()
}

func (m *Model) Blur() {
	m.Focused = 0
	m.dialog.Blur()
}

// SetText
//
// 显示我的回复信息
// id means vxid
func (m *Model) SetText(id, nick, s string) {
	// 消息人是当前对话窗口联系人
	if id != "" && id != m.dialog.user.wid {
		return
	}

	var b strings.Builder
	if nick != "" {
		b.WriteString("(")
		b.WriteString(nick)
		b.WriteString("): ")
	}
	b.WriteString(s)
	b.WriteString("\n")

	txt := b.String()
	m.textarea.InsertString(txt)
	record.SetTxtById(id, txt)
	//TODO cursor down didn't work ?
	//m.textarea.CursorEnd()
	//m.textarea.CursorDown()

	// textarea can't cursor down
	if m.textarea.Line() >= m.textarea.Height() {
		m.Reset()
	}
}

func (m *Model) Reset() {
	m.textarea.Reset()
}
