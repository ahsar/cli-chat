// user input box
package message

import (
	"log"
	"strconv"

	"github.com/ahsar/cli-chat/internal/chat"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	esc, send key.Binding
}

type user struct {
	id  int
	wid string // vx id
}

type DialogModel struct {
	textarea textarea.Model
	keymap   keymap
	user     user // current dialog user
}

func NewDialogModel() *DialogModel {
	t := textarea.New()
	t.Prompt = ""
	t.ShowLineNumbers = true
	t.Placeholder = "Press Ctrl+enter to send"
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

	return &DialogModel{
		textarea: t,
		keymap: keymap{
			send: key.NewBinding(
				key.WithKeys("ctrl+s"),
			),
			esc: key.NewBinding(
				key.WithKeys("ctrl+q"),
			),
		},
	}
}

func (m DialogModel) Init() (t tea.Cmd) {
	return
}

func (m *DialogModel) Update(msg tea.Msg) (*DialogModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.send):
			if txt := m.textarea.Value(); txt != "" {
				// 1. send content to vx
				chat.TalkToId(m.user.id, txt)

				// 2. send content to message panel
				Msg.SetText("", "我", txt)

				// 3. clear dialog panel input
				m.ClearInput()
			}

		case key.Matches(msg, m.keymap.esc):
			log.Println("exit dialog panel")
			m.Blur()
		}
	}
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

func (m *DialogModel) ClearInput() {
	m.textarea.SetValue("")
}
func (m *DialogModel) SetUser(id string) {
	i, _ := strconv.Atoi(id)
	m.user.id = i
	m.user.wid = chat.FriendById(i).ID()
}
