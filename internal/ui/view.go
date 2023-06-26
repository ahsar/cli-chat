package ui

import (
	"log"

	"github.com/ahsar/cli-chat/internal/chat"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/table"
)

func (m *model) sizeInput() {
	log.Printf(
		"resize h:%d w:%d",
		m.height,
		m.width,
	)

	m.rencent.SetSize(
		m.width/3,
		m.height-constant.HelpHeight-3)

	m.message.SetSize(m.width, m.height)

	m.contacts.SetSize(
		m.width/3,
		m.height-constant.HelpHeight-3)
}

func (m *model) blur() {
	m.contacts.Blur()
	m.rencent.Blur()
}

func (m *model) focus() {
	if m.current == constant.ContactPanel {
		m.current = constant.RencentPanel
		m.contacts.Blur()
		m.rencent.Focus()
	} else {
		m.current = constant.ContactPanel
		m.rencent.Blur()
		m.contacts.Focus()
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
