package ui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	cursorLineStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("57")).
			Foreground(lipgloss.Color("230"))

	placeholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))

	endOfBufferStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("235"))

	focusedPlaceholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("99"))

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color("238"))

	blurredBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder())
)

func (m model) sizeInput() {
	l := len(m.inputs)
	for i := range m.inputs {

		switch i {
		// 最近联系人
		case 0:
			m.inputs[i].SetWidth(m.width / l)
			m.inputs[i].SetHeight(m.height - helpHeight - 2)
		// 聊天窗口
		case 1:
			m.inputs[i].SetWidth(m.width / l)
			m.inputs[i].SetHeight(m.height/2 - helpHeight - 2)
		// 对话框
		case 2:
			m.inputs[i].SetWidth(m.width / l)
			m.inputs[i].SetHeight(m.height/2 - helpHeight - 1)
		// 通讯录
		case 3:
			m.inputs[i].SetWidth(m.width / l)
			m.inputs[i].SetHeight(m.height - helpHeight)
		}
	}
}

func newTextarea(i int) textarea.Model {
	t := textarea.New()
	t.Prompt = ""
	//t.Placeholder = "Type something"
	t.ShowLineNumbers = false
	t.Cursor.Style = cursorStyle
	t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	t.BlurredStyle.Placeholder = placeholderStyle
	t.FocusedStyle.CursorLine = cursorLineStyle
	t.FocusedStyle.Base = focusedBorderStyle
	t.BlurredStyle.Base = blurredBorderStyle
	t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	//t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	t.SetValue("\n\n\n" + strconv.Itoa(i+1))
	return t
}
