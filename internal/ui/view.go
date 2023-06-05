package ui

import (
	"github.com/charmbracelet/bubbles/textarea"
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
	//t.ShowLineNumbers = false
	//t.Cursor.Style = cursorStyle
	//t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	//t.BlurredStyle.Placeholder = placeholderStyle
	//t.FocusedStyle.CursorLine = cursorLineStyle
	//t.FocusedStyle.Base = focusedBorderStyle
	//t.BlurredStyle.Base = blurredBorderStyle
	//t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	//t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	////t.KeyMap.DeleteWordBackward.SetEnabled(false)
	//t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	//t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	//t.Blur()
	//t.SetValue("\n\n\n" + strconv.Itoa(i+1))
	return t
}
