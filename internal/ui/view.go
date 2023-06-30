package ui

import (
	"log"

	"github.com/ahsar/cli-chat/internal/chat"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/table"
	"github.com/eatmoreapple/openwechat"
)

func (m *model) sizeInput() {
	m.rencent.SetSize(
		m.width/3,
		m.height-constant.HelpHeight-3)

	m.message.SetSize(m.width, m.height)

	m.contacts.SetSize(
		m.width/3,
		m.height-constant.HelpHeight-3)
}

func (m *model) exit() {
	m.contacts.Blur()
	m.rencent.Blur()
	m.message.Blur()
}

func (m *model) getCurrent() (b byte) {
	return m.contacts.Focused |
		m.message.Focused |
		m.rencent.Focused
}

func (m *model) focusInTurn() {
	i := m.getCurrent()
	switch i {
	case i & constant.ContactPanel:
		m.contacts.Blur()
		m.rencent.Focus()
	case i & constant.RencentPanel:
		m.contacts.Focus()
		m.rencent.Blur()
	default:
		m.contacts.Focus()
		m.rencent.Blur()
		m.message.Blur()
	}
}

// 设定通讯录
func (m *model) SetContacts() {
	frList := chat.Friends()
	l := len(frList)
	if l <= 0 {
		return
	}

	var rows = make([]table.Row, 0, l)
	for _, v := range frList {
		rows = append(rows, v)
	}
	m.contacts.SetRow(rows)
}

func (m *model) onMsg(msg *openwechat.Message) {
	u, err := msg.Sender()
	if err != nil {
		log.Println("err get msg send", err)
	}
	content := msg.Content
	if !msg.IsText() {
		content = "[" + msg.MsgType.String() + "], 暂不支持在终端查看, 请前往手机查看"
	}

	m.message.SetText(u.ID(), chat.GetName(u), content)
}
