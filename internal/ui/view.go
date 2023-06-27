package ui

import (
	"log"

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
	}
}

// 设定通讯录
func (m *model) SetContacts() {
	//TODO
	frList := [][]string{
		{"1", "x"},
		{"2", "y"},
		{"3", "z"},
		{"4", "t"},
	}
	//frList := chat.Friends()
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
